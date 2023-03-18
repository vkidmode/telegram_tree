package tree

import (
	"context"
	"fmt"
)

type nextGeneratorFunc func(ctx context.Context, chatID int64) (NodesList, error)
type processorFunc func(ctx context.Context, chatID int64, callBack string) error

func NewNodesHandler(ctx context.Context, template NodesList, chatID int64, defaultMessage string) (*NodesHandler, error) {
	var handler = &NodesHandler{
		ctx:            ctx,
		chatID:         chatID,
		defaultMessage: defaultMessage,
		templateTree:   template,
	}
	return handler, nil
}

type NodesHandler struct {
	defaultMessage string
	chatID         int64
	ctx            context.Context
	templateTree   NodesList
}

func (n *NodesHandler) GetFilledNodesList() (NodesList, error) {
	var node NodesList
	node = n.templateTree

	for i := range node {
		if err := checkAndFillSingleNode(n.ctx, node[i], n.chatID, n.defaultMessage, i); err != nil {
			return nil, fmt.Errorf("checking node: %w", err)
		}
	}
	return node, nil
}

func checkAndFillSingleNode(ctx context.Context, node *Node, chatID int64, defaultMessage string, id int) (err error) {
	if node == nil {
		return nil
	}

	if err = node.checkValidity(); err != nil {
		return err
	}
	if node.id, err = convertNumberToSymbol(id); err != nil {
		return err
	}
	node.setDefaultMessageIfNeed(defaultMessage)
	node.ctx = ctx
	node.chatID = chatID

	for i := range node.NextNodes {
		if err = checkAndFillSingleNode(ctx, node.NextNodes[i], chatID, defaultMessage, i); err != nil {
			return err
		}
	}
	return nil
}

func (n *NodesHandler) GetNodeByCallback(callback string) (*Node, error) {
	symbolsList, err := parseCallback(callback)
	if err != nil {
		return nil, err
	}
	numbersList, err := symbolsList.toNumbers()
	if err != nil {
		return nil, err
	}

	nodes, err := n.GetFilledNodesList()
	if err != nil {
		return nil, err
	}

	var currentNode *Node

	for i := range numbersList {
		if i == 0 {
			currentNode = nodes[numbersList[i]]
			continue
		}
		currentNode = currentNode.NextNodes[numbersList[i]]
	}
	return currentNode, nil
}
