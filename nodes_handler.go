package telegram_tree

import (
	"context"
	"fmt"
	"strings"
)

type ProcessorFunc func(ctx context.Context, meta Meta) ([]Node, error)

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

func (n *NodesHandler) GetNode(ctx context.Context, meta Meta) (Node, error) {
	if n.templateTree == nil {
		return nil, nil
	}

	if _, err := parseCallback(meta.GetCallback()); err != nil {
		return nil, err
	}

	elements, err := getElementsFromCallback(meta.GetCallback())
	if err != nil {
		return nil, err
	}

	currentNode := &node{nextNodes: n.templateTree}

	for i := range elements {
		symbol, err := extractSymbolFromElem(elements[i])
		if err != nil {
			return nil, err
		}

		meta.SetupCallback(strings.Join(elements[:i+1], callbackDivider))
		if symbol == CallBackSkip {
			if err = currentNode.jumpToNode(currentNode.skip); err != nil {
				return nil, fmt.Errorf("skippting callback: %w", err)
			}
			currentNode.callback = meta.GetCallback()
			if err = currentNode.nextNodes.setupCallBacks(currentNode.callback); err != nil {
				return nil, fmt.Errorf("setting up callback")
			}
		} else {
			number, err := convertSymbolToNum(symbol)
			if err != nil {
				return nil, fmt.Errorf("error converting symbol to number")
			}
			if err = currentNode.fillNextNodes(ctx, meta); err != nil {
				return nil, fmt.Errorf("getting next nodes for non root node: %v", err)
			}
			if err = currentNode.jumpToChild(number); err != nil {
				return nil, err
			}
		}

		meta.SetIsMiddle(true)
		if i == (len(elements) - 1) {
			meta.SetIsMiddle(false)
		}

		if err = currentNode.fillNextNodes(ctx, meta); err != nil {
			return nil, fmt.Errorf("getting next nodes for non root node: %v", err)
		}
	}
	currentNode.setDefaultMessageIfNeed(n.defaultMessage)
	currentNode.callback = meta.GetCallback()
	return currentNode, nil
}

func ExtractPayload(callBack string) (map[string]string, error) {
	return extractPayloadFromCallback(callBack)
}

func (n *NodesHandler) checkSingleNode(node Node) error {
	if node == nil {
		return nil
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
