package tree

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
	Processor          processorFunc
	NextNodesGenerator nextGeneratorFunc
	chatID             int64
	ctx                context.Context
	id                 string
}

type NodesList []*Node

func (n NodesList) convertToList() (resp []string) {
	for i := range n {
		resp = append(resp, n[i].GetButtonText())
	}
	return resp
}

func (n *Node) GetMessage() string    { return n.Message }
func (n *Node) IsBarHidden() bool     { return n.HideBar }
func (n *Node) GetButtonText() string { return n.HumanText }
func (n *Node) GetNextButtonsText() ([]string, error) {
	var err error
	if len(n.NextNodes) > 0 {
		return n.NextNodes.convertToList(), nil
	}
	if n.NextNodesGenerator != nil {
		n.NextNodes, err = n.NextNodesGenerator(n.ctx, n.chatID)
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
