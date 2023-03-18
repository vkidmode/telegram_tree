package tree

import (
	"context"
	"fmt"
)

type nextGeneratorFunc func(ctx context.Context, chatID int64) (NodesList, error)
type processorFunc func(ctx context.Context, chatID int64, callBack string) error

type NodesHandler struct {
	defaultMessage string
	templateTree   NodesList
}

func NewNodesHandler(template NodesList, defaultMessage string) (*NodesHandler, error) {
	var handler = &NodesHandler{
		defaultMessage: defaultMessage,
		templateTree:   template,
	}
	if err := handler.checkTemplate(); err != nil {
		return nil, err
	}
	return handler, nil
}

func (n *NodesHandler) checkTemplate() error {
	if n.templateTree == nil {
		return fmt.Errorf("template is null")
	}
	for i := range n.templateTree {
		if err := n.checkSingleNode(n.templateTree[i]); err != nil {
			return err
		}
	}
	return nil
}

func (n *NodesHandler) checkSingleNode(node *Node) error {
	if node == nil {
		return fmt.Errorf("null node")
	}
	if err := node.checkValidity(); err != nil {
		return err
	}
	for i := range node.NextNodes {
		if err := n.checkSingleNode(node.NextNodes[i]); err != nil {
			return err
		}
	}
	return nil
}

func (n *NodesHandler) GetNodeByCallback(ctx context.Context, chatID int64, callback string) (*Node, error) {
	symbolsList, err := parseCallback(callback)
	if err != nil {
		return nil, err
	}
	numbersList, err := symbolsList.toNumbers()
	if err != nil {
		return nil, err
	}

	if n.templateTree == nil {
		return nil, nil
	}

	var currentNode = &Node{
		NextNodes: n.templateTree,
	}

	for i := range numbersList {
		if i == 0 {
			if err = currentNode.jumpToChild(numbersList[i]); err != nil {
				return nil, fmt.Errorf("invalid callback")
			}
			continue
		}
		if err = currentNode.fillNextNodes(ctx, chatID); err != nil {
			return nil, err
		}
		if err = currentNode.jumpToChild(numbersList[i]); err != nil {
			return nil, err
		}
	}
	currentNode.setDefaultMessageIfNeed(n.defaultMessage)
	return currentNode, nil
}
