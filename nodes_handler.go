package telegram_tree

import (
	"context"
	"fmt"
)

type NextGeneratorFunc func(ctx context.Context, chatID int64) ([]Node, error)
type ProcessorFunc func(ctx context.Context, messageID int, chatID int64, callBack string) error

type NodesHandler struct {
	defaultMessage string
	templateTree   nodesList
}

func NewNodesHandler(template nodesList, defaultMessage string) (*NodesHandler, error) {
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

func (n *NodesHandler) checkSingleNode(node Node) error {
	if node == nil {
		return fmt.Errorf("null node")
	}
	if err := node.checkValidity(); err != nil {
		return err
	}
	for i := range node.getNextNodes() {
		if err := n.checkSingleNode(node.getNextNodes()[i]); err != nil {
			return err
		}
	}
	return nil
}

func (n *NodesHandler) GetNodeByCallback(ctx context.Context, chatID int64, callback string) (Node, error) {
	symbolsList, err := parseCallback(callback)
	if err != nil {
		return nil, err
	}

	if n.templateTree == nil {
		return nil, nil
	}

	var currentNode = &node{
		nextNodes: n.templateTree,
	}

	for i := range symbolsList {
		if symbolsList[i] == CallBackSkip {
			if currentNode.skip == nil {
				return nil, fmt.Errorf("invalid callback")
			}
			currentNode.jumpToNode(currentNode.skip)
			continue
		}
		number, err := convertSymbolToNum(symbolsList[i])
		if err != nil {
			return nil, fmt.Errorf("error converting symbol to number")
		}

		if i == 0 {
			if err = currentNode.jumpToChild(number); err != nil {
				return nil, fmt.Errorf("invalid callback")
			}
			continue
		}
		if err = currentNode.fillNextNodes(ctx, chatID); err != nil {
			return nil, err
		}
		if err = currentNode.jumpToChild(number); err != nil {
			return nil, err
		}
	}
	currentNode.setDefaultMessageIfNeed(n.defaultMessage)
	currentNode.callback = callback
	return currentNode, nil
}
