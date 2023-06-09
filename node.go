package telegram_tree

import (
	"context"
	"fmt"
	"strings"
)

type node struct {
	payload            Payload
	skip               Node
	processor          ProcessorFunc
	nextNodesGenerator NextGeneratorFunc
	nextNodes          nodesList
	callback           string
	telegramOptions    TelegramOptions
}

func (n *node) toInterface() Node {
	return n
}

func (n *node) GetTelegramOptions() TelegramOptions      { return n.telegramOptions }
func (n *node) GetProcessor() ProcessorFunc              { return n.processor }
func (n *node) GetNextNodesGenerator() NextGeneratorFunc { return n.nextNodesGenerator }
func (n *node) ExtractPayload() (map[string]string, error) {
	return extractPayloadFromCallback(n.callback)
}

func (n *node) setTelegramOptions(in TelegramOptions) { n.telegramOptions = in }
func (n *node) setProcessor(in ProcessorFunc)         { n.processor = in }
func (n *node) setNextGenerator(in NextGeneratorFunc) { n.nextNodesGenerator = in }
func (n *node) setNextNodes(in []Node)                { n.nextNodes = in }
func (n *node) getNextNodes() []Node                  { return n.nextNodes }
func (n *node) setCallback(in string)                 { n.callback = in }
func (n *node) GetPayload() Payload                   { return n.payload }
func (n *node) setSkipper(in Node)                    { n.skip = in }
func (n *node) setPayload(in Payload)                 { n.payload = in }

type Node interface {
	GetProcessor() ProcessorFunc
	GetNextNodesGenerator() NextGeneratorFunc
	GetNextNodes(ctx context.Context, chatID int64) ([]Node, error)
	GetCallback() string
	GetCallbackBack() (string, error)
	GetCallbackSkip() (string, error)
	GetPayload() Payload
	ExtractPayload() (map[string]string, error)
	GetTelegramOptions() TelegramOptions

	setTelegramOptions(in TelegramOptions)
	setProcessor(ProcessorFunc)
	setNextGenerator(NextGeneratorFunc)
	setNextNodes([]Node)
	setPayload(Payload)
	setCallback(string)
	getInternalStruct() *node
	checkValidity() error
	getNextNodes() []Node
	setSkipper(in Node)
}

func NewNode(
	telegramOptions TelegramOptions,
	payloadItem Payload,
	processor ProcessorFunc,
	nextNodesGenerator NextGeneratorFunc,
	skipNodeGenerator Node,
) Node {
	var nodeItem = &node{}
	nodeInterface := nodeItem.toInterface()
	nodeInterface.setTelegramOptions(telegramOptions)
	nodeInterface.setProcessor(processor)
	nodeInterface.setNextGenerator(nextNodesGenerator)
	nodeInterface.setPayload(payloadItem)
	nodeInterface.setSkipper(skipNodeGenerator)
	return nodeInterface
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

func (n *node) getInternalStruct() *node { return n }

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
	n.jumpToNode(n.nextNodes[in])
	return nil
}

func (n *node) jumpToNode(node Node) {
	internalStruct := node.getInternalStruct()
	*n = *internalStruct
}

func (n *node) GetNextNodes(ctx context.Context, chatID int64) ([]Node, error) {
	var err error
	if len(n.nextNodes) == 0 {
		if n.nextNodesGenerator != nil {
			if n.nextNodes, err = n.nextNodesGenerator(ctx, chatID); err != nil {
				return nil, err
			}
		}
	}
	if err = n.nextNodes.setupCallBacks(n.callback); err != nil {
		return nil, err
	}
	return n.nextNodes, nil
}

func (n *node) setDefaultMessageIfNeed(defMsg string) {
	if n.telegramOptions.GetMessage() == "" {
		n.telegramOptions.setDefaultMessage(defMsg)
	}
}

func (n *node) checkValidity() error {
	if n.GetTelegramOptions().GetHumanText() == "" {
		return fmt.Errorf("each node should have human text")
	}
	if n.GetTelegramOptions().GetHideBar() && n.skip != nil {
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
