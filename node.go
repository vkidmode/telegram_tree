package telegram_tree

import (
	"context"
	"fmt"
)

type Node struct {
	Message            string
	HumanText          string
	HideBar            bool
	Payload            string
	Skip               *Node
	NextNodes          NodesList
	Processor          ProcessorFunc
	NextNodesGenerator NextGeneratorFunc
}

func (n *Node) fillNextNodes(ctx context.Context, chatID int64) (err error) {
	if n.NextNodes == nil {
		if n.NextNodesGenerator != nil {
			if n.NextNodes, err = n.NextNodesGenerator(ctx, chatID); err != nil {
				return err
			}
		}
	}
	if n.NextNodes == nil {
		return fmt.Errorf("next nodes not available")
	}
	return nil
}

func (n *Node) jumpToChild(in int) error {
	if in < 0 {
		return fmt.Errorf("invalid number")
	}
	if in > len(n.NextNodes)-1 {
		return fmt.Errorf("invalid number")
	}
	if n.NextNodes[in] == nil {
		return fmt.Errorf("child is null")
	}
	*n = *n.NextNodes[in]
	return nil
}

func (n *Node) GetMessage() string    { return n.Message }
func (n *Node) IsBarHidden() bool     { return n.HideBar }
func (n *Node) GetButtonText() string { return n.HumanText }
func (n *Node) GetNextButtonsText(ctx context.Context, chatID int64) ([]string, error) {
	var err error
	if len(n.NextNodes) > 0 {
		return n.NextNodes.convertToList(), nil
	}
	if n.NextNodesGenerator != nil {
		n.NextNodes, err = n.NextNodesGenerator(ctx, chatID)
		if err != nil {
			return nil, err
		}
		return n.NextNodes.convertToList(), nil
	}
	return nil, nil
}

func (n *Node) setDefaultMessageIfNeed(defMsg string) {
	if n.Message == "" {
		n.Message = defMsg
	}
}

func (n *Node) checkValidity() error {
	if n.HumanText == "" {
		return fmt.Errorf("each node should have human text")
	}
	if n.HideBar && n.Skip != nil {
		return fmt.Errorf("skip function works only with not hidden bar")
	}
	if n.Processor != nil && n.NextNodesGenerator != nil {
		return fmt.Errorf("unable to have processor and nextNodesGenerator in same time")
	}
	if n.NextNodesGenerator != nil && len(n.NextNodes) > 0 {
		return fmt.Errorf("unable to have nextNodes and nextNodesGenerator in same time")
	}
	return nil
}

type NodesList []*Node

func (n NodesList) convertToList() (resp []string) {
	for i := range n {
		resp = append(resp, n[i].GetButtonText())
	}
	return resp
}
