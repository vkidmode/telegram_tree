package telegram_tree

import (
	"context"
	"fmt"
	"strings"
)

type NodeOpt func(v *node)

type node struct {
	payload   Payload
	skip      Node
	processor ProcessorFunc
	telegram  Telegram
	nextNodes nodesList
	callback  string
}

type Node interface {
	GetProcessor() ProcessorFunc
	GetCallback() string
	GetCallbackBack() (string, error)
	GetCallbackSkip() (string, error)
	GetTelegram() Telegram
	GetChildren() []Node

	setCallback(string)
	getInternalStruct() *node
	checkValidity() error
	getNextNodes() []Node
	getPayload() Payload
}

func NewNode(options ...NodeOpt) Node {
	var nodeItem = node{}
	for _, opt := range options {
		opt(&nodeItem)
	}
	return &nodeItem
}

func WithSkipper(in Node) NodeOpt {
	return func(v *node) {
		v.skip = in
	}
}

func WithPayload(payload Payload) NodeOpt {
	return func(v *node) {
		v.payload = payload
	}
}

func WithProc(proc ProcessorFunc) NodeOpt {
	return func(v *node) {
		v.processor = proc
	}
}

func WithTg(tg Telegram) NodeOpt {
	return func(v *node) {
		v.telegram = tg
	}
}

func (n *node) GetChildren() []Node         { return n.nextNodes }
func (n *node) GetTelegram() Telegram       { return n.telegram }
func (n *node) GetProcessor() ProcessorFunc { return n.processor }
func (n *node) GetCallback() string         { return n.callback }

func (n *node) GetCallbackBack() (string, error) {
	currentCallback := n.GetCallback()

	currentCallbackElements, err := parseCallback(currentCallback) // dsfgsdfg
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
	if _, err := parseCallback(currentCallback); err != nil {
		return "", err
	}
	callBackParts := strings.Split(currentCallback, callbackDivider)
	callBackParts = append(callBackParts, CallBackSkip)
	return strings.Join(callBackParts, callbackDivider), nil
}

func (n *node) fillNextNodes(ctx context.Context, meta Meta) error {
	var err error

	if len(n.nextNodes) == 0 {
		if n.processor != nil {
			if n.nextNodes, err = n.processor(ctx, meta); err != nil {
				return err
			}
		}
	}

	if err = n.nextNodes.setupCallBacks(n.callback); err != nil {
		return err
	}
	return nil
}

func (n *node) jumpToChild(in int) error {
	if in < 0 {
		return fmt.Errorf("invalid number cannot use negative numbers")
	}
	if in > len(n.nextNodes)-1 {
		return fmt.Errorf("invalid number too big %d, max is %d name is %s", in, len(n.nextNodes)-1, n.telegram.GetTabTxt())
	}
	return n.jumpToNode(n.nextNodes[in])
}

func (n *node) jumpToNode(node Node) error {
	if node == nil {
		return fmt.Errorf("cannot jump to null node")
	}
	internalStruct := node.getInternalStruct()
	*n = *internalStruct
	return nil
}

func (n *node) setDefaultMessageIfNeed(defMsg string) {
	if n.telegram.GetMessage() == "" {
		n.telegram.setDefaultMessage(defMsg)
	}
}

func (n *node) checkValidity() error {
	if n.GetTelegram().GetTabTxt() == "" {
		return fmt.Errorf("each node should have human text")
	}
	if n.GetTelegram().GetHideBar() && n.skip != nil {
		return fmt.Errorf("skip function works only with not hidden bar")
	}
	return nil
}

func (n *node) getNextNodes() []Node { return n.nextNodes }

func (n *node) setCallback(in string) { n.callback = in }

func (n *node) getPayload() Payload { return n.payload }

func (n *node) getInternalStruct() *node { return n }
