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
	skip               NextGeneratorFunc
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
func (n *node) GetSkip() NextGeneratorFunc               { return n.skip }
func (n *node) GetProcessor() ProcessorFunc              { return n.processor }
func (n *node) GetNextNodesGenerator() NextGeneratorFunc { return n.nextNodesGenerator }

func (n *node) setMessage(in string)                  { n.message = in }
func (n *node) setHumanText(in string)                { n.humanText = in }
func (n *node) setHideBar(in bool)                    { n.hideBar = in }
func (n *node) setProcessor(in ProcessorFunc)         { n.processor = in }
func (n *node) setNextGenerator(in NextGeneratorFunc) { n.nextNodesGenerator = in }
func (n *node) setNextNodes(in []Node)                { n.nextNodes = in }
func (n *node) getNextNodes() []Node                  { return n.nextNodes }
func (n *node) setCallback(in string)                 { n.callback = in }
func (n *node) GetPayload() Payload                   { return &n.payload }
func (n *node) setSkipper(in NextGeneratorFunc)       { n.skip = in }
func (n *node) setPayload(in Payload) {
	if in != nil {
		n.payload.value = in.GetValue()
		n.payload.key = in.GetKey()
	}
}

type Node interface {
	GetMessage() string
	GetHumanText() string
	GetHideBar() bool
	GetSkip() NextGeneratorFunc
	GetProcessor() ProcessorFunc
	GetNextNodesGenerator() NextGeneratorFunc
	GetNextNodes(ctx context.Context, chatID int64) ([]Node, error)
	GetCallback() string
	GetCallbackBack() (string, error)
	GetCallbackSkip() (string, error)
	GetPayload() Payload

	setMessage(string)
	setHumanText(string)
	setHideBar(bool)
	setProcessor(ProcessorFunc)
	setNextGenerator(NextGeneratorFunc)
	setNextNodes([]Node)
	setPayload(Payload)
	setCallback(string)
	getInternalStruct() *node
	checkValidity() error
	getNextNodes() []Node
	setSkipper(in NextGeneratorFunc)
}

func NewNode(
	message string,
	humanText string,
	payloadItem Payload,
	hideBar bool,
	processor ProcessorFunc,
	nextNodesGenerator NextGeneratorFunc,
	skipNodeGenerator NextGeneratorFunc,
	nextNodes ...Node,
) Node {
	var nodeItem = &node{}
	nodeInterface := nodeItem.toInterface()
	nodeInterface.setMessage(message)
	nodeInterface.setHumanText(humanText)
	nodeInterface.setHideBar(hideBar)
	nodeInterface.setProcessor(processor)
	nodeInterface.setNextGenerator(nextNodesGenerator)
	nodeInterface.setPayload(payloadItem)
	nodeInterface.setNextNodes(nextNodes)
	nodeInterface.setSkipper(skipNodeGenerator)
	return nodeInterface
}

//func convertInterfaceToNodes(input []Node) nodesList {
//	resp := make(nodesList, len(input))
//	for i := range input {
//		if val, ok := input[i].(*node); ok {
//			resp[i] = val
//		} else {
//			return nil
//		}
//	}
//	return resp
//}

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
	internalStruct := n.nextNodes[in].getInternalStruct()
	*n = *internalStruct
	return nil
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
