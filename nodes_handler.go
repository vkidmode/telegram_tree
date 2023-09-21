package telegram_tree

import (
	"context"
	"fmt"
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

func (n *NodesHandler) GetNode(ctx context.Context, meta Meta) (Node, error) {
	if n.templateTree == nil {
		return nil, nil
	}

	symbolsList, err := parseCallback(meta.GetCallback())
	if err != nil {
		return nil, err
	}

	currentNode := &node{
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
			nullChild, err := currentNode.jumpToChild(number)
			if err != nil {
				return nil, fmt.Errorf("invalid callback")
			}
			if nullChild {
				return nil, nil
			}
			continue
		}

		if err = currentNode.fillNextNodes(ctx, meta); err != nil {
			return nil, err
		}

		nullChild, err := currentNode.jumpToChild(number)
		if err != nil {
			return nil, err
		}

		if nullChild {
			return nil, nil
		}
	}
	currentNode.setDefaultMessageIfNeed(n.defaultMessage)
	currentNode.callback = meta.GetCallback()
	return currentNode, nil
}

func ExtractPayload(callBack string) (map[string]string, error) {
	return extractPayloadFromCallback(callBack)
}
