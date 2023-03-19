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
	Processor          ProcessorFunc
	NextNodesGenerator NextGeneratorFunc
	callback           string
	nextNodes          NodesList
}

func (n *Node) fillNextNodes(ctx context.Context, chatID int64) (err error) {
	if n.nextNodes == nil {
		if n.NextNodesGenerator != nil {
			if n.nextNodes, err = n.NextNodesGenerator(ctx, chatID); err != nil {
				return err
			}
		}
	}
	if n.nextNodes == nil {
		return fmt.Errorf("next nodes not available")
	}
	return nil
}

func (n *Node) GetCallBack() string {
	return n.callback
}

func (n *Node) jumpToChild(in int) error {
	if in < 0 {
		return fmt.Errorf("invalid number")
	}
	if in > len(n.nextNodes)-1 {
		return fmt.Errorf("invalid number")
	}
	if n.nextNodes[in] == nil {
		return fmt.Errorf("child is null")
	}
	*n = *n.nextNodes[in]
	return nil
}

func (n *Node) GetMessage() string    { return n.Message }
func (n *Node) IsBarHidden() bool     { return n.HideBar }
func (n *Node) GetButtonText() string { return n.HumanText }

func (n *Node) GetNextNodes(ctx context.Context, chatID int64) (NodesList, error) {
	var err error
	if len(n.nextNodes) == 0 {
		if n.NextNodesGenerator != nil {
			n.nextNodes, err = n.NextNodesGenerator(ctx, chatID)
			if err != nil {
				return nil, err
			}
		}
	}
	if err = n.nextNodes.setupCallBacks(n.callback); err != nil {
		return nil, err
	}
	return n.nextNodes, nil
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
	if n.NextNodesGenerator != nil && len(n.nextNodes) > 0 {
		return fmt.Errorf("unable to have nextNodes and nextNodesGenerator in same time")
	}
	return nil
}

type NodesList []*Node

func (n NodesList) setupCallBacks(callback string) error {
	for i := range n {
		newCallback, err := incrementCallback(callback, i)
		if err != nil {
			return err
		}
		n[i].callback = newCallback
	}
	return nil
}
