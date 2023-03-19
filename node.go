package telegram_tree

import (
	"context"
	"fmt"
	"strings"
)

type node struct {
	message            string
	humanText          string
	hideBar            bool
	payload            payload
	skip               ProcessorFunc
	processor          ProcessorFunc
	nextNodesGenerator NextGeneratorFunc
	nextNodes          nodesList
	callback           string
}

func (n *node) toInterface() Node {
	return n
}

func (n *node) GetMessage() string                       { return n.message }
func (n *node) GetHumanText() string                     { return n.humanText }
func (n *node) GetHideBar() bool                         { return n.hideBar }
func (n *node) GetSkip() ProcessorFunc                   { return n.skip }
func (n *node) GetProcessor() ProcessorFunc              { return n.processor }
func (n *node) GetNextNodesGenerator() NextGeneratorFunc { return n.nextNodesGenerator }

type Node interface {
	GetMessage() string
	GetHumanText() string
	GetHideBar() bool
	GetSkip() ProcessorFunc
	GetProcessor() ProcessorFunc
	GetNextNodesGenerator() NextGeneratorFunc
	GetNextNodes(ctx context.Context, chatID int64) ([]Node, error)
	GetCallback() string
	GetCallbackBack() (string, error)
	GetCallbackSkip() (string, error)
}

func NewNode(
	message string,
	humanText string,
	payloadItem Payload,
	hideBar bool,
	nextNodes ...Node,
) Node {
	nodeItem := &node{
		message:   message,
		humanText: humanText,
		hideBar:   hideBar,
	}
	if val, ok := payloadItem.(*payload); ok {
		if val != nil {
			nodeItem.payload = *val
		} else {
			return nil
		}
	} else {
		return nil
	}

	nodeItem.nextNodes = make(nodesList, len(nextNodes))
	for i := range nextNodes {
		if val, ok := nextNodes[i].(*node); ok {
			nodeItem.nextNodes[i] = val
		} else {
			return nil
		}
	}

	return nodeItem
}

func (n *node) fillNextNodes(ctx context.Context, chatID int64) (err error) {
	if n.nextNodes == nil {
		if n.nextNodesGenerator != nil {
			if n.nextNodes, err = n.nextNodesGenerator(ctx, chatID); err != nil {
				return err
			}
		}
	}
	if n.nextNodes == nil {
		return fmt.Errorf("next nodes not available")
	}
	return nil
}

func (n *node) GetCallback() string { return n.callback }

func (n *node) GetCallbackBack() (string, error) {
	currentCallback := n.GetCallback()
	currentCallbackElements, err := parseCallback(currentCallback)
	if err != nil {
		return "", err
	}
	if len(currentCallbackElements) < 2 {
		return "", nil
	}
	callBackParts := strings.Split(currentCallback, callbackDivider)
	callBackParts = callBackParts[:len(callBackParts)-1]
	return strings.Join(callBackParts, callbackDivider), nil
}

func (n *node) GetCallbackSkip() (string, error) {
	if n.skip == nil {
		return "", nil
	}
	currentCallback := n.GetCallback()
	_, err := parseCallback(currentCallback)
	if err != nil {
		return "", err
	}
	callBackParts := strings.Split(currentCallback, callbackDivider)
	callBackParts = append(callBackParts, CallBackSkip)
	return strings.Join(callBackParts, callbackDivider), nil
}

func (n *node) jumpToChild(in int) error {
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

func (n *node) GetNextNodes(ctx context.Context, chatID int64) ([]Node, error) {
	var err error
	if len(n.nextNodes) == 0 {
		if n.nextNodesGenerator != nil {
			n.nextNodes, err = n.nextNodesGenerator(ctx, chatID)
			if err != nil {
				return nil, err
			}
		}
	}
	if err = n.nextNodes.setupCallBacks(n.callback); err != nil {
		return nil, err
	}
	return n.nextNodes.toInterface(), nil
}

func (n *node) setDefaultMessageIfNeed(defMsg string) {
	if n.message == "" {
		n.message = defMsg
	}
}

func (n *node) checkValidity() error {
	if n.humanText == "" {
		return fmt.Errorf("each node should have human text")
	}
	if n.hideBar && n.skip != nil {
		return fmt.Errorf("skip function works only with not hidden bar")
	}
	if n.processor != nil && n.nextNodesGenerator != nil {
		return fmt.Errorf("unable to have processor and nextNodesGenerator in same time")
	}
	if n.nextNodesGenerator != nil && len(n.nextNodes) > 0 {
		return fmt.Errorf("unable to have nextNodes and nextNodesGenerator in same time")
	}
	return nil
}
