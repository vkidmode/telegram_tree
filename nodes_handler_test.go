package telegram_tree

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkCallbackElement(t *testing.T) {
	for i, tCase := range []struct {
		element string
		valid   bool
	}{
		{
			element: "kaka",
			valid:   false,
		},
		{
			element: "a",
			valid:   true,
		},
		{
			element: "",
			valid:   false,
		},
		{
			element: "A",
			valid:   false,
		},
		{
			element: "z(dasgljsdfgn2!)",
			valid:   true,
		},
		{
			element: "z(aaaa)z",
			valid:   false,
		},
		{
			element: "z()",
			valid:   false,
		},
		{
			element: "()",
			valid:   false,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			valid, err := checkCallbackElement(tCase.element)
			if err != nil {
				t.Errorf("error in checkCallbackElement: %v", err)
			}
			assert.Equal(t, tCase.valid, valid)
		})
	}
}

func Test_extractSymbolFromElem(t *testing.T) {
	for i, tCase := range []struct {
		element   string
		haveError bool
		symbol    string
	}{
		{
			element: "z",
			symbol:  "z",
		},
		{
			element:   "z))",
			haveError: true,
		},
		{
			element: "a(ksksk)",
			symbol:  "a",
		},
		{
			element:   "aa(ksksk)",
			haveError: true,
		},
		{
			element:   "",
			haveError: true,
		},
		{
			element:   " ",
			haveError: true,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			symbol, err := extractSymbolFromElem(tCase.element)
			if err != nil {
				assert.Equal(t, tCase.haveError, true)
			} else {
				assert.Equal(t, tCase.symbol, symbol)
			}
		})
	}
}

func Test_parseCallback(t *testing.T) {
	for i, tCase := range []struct {
		callback  string
		haveError bool
		resp      []string
	}{
		{
			callback:  "",
			haveError: true,
		},
		{
			callback: "a",
			resp:     []string{"a"},
		},
		{
			callback:  "a-----",
			haveError: true,
		},
		{
			callback:  "a--b",
			haveError: true,
		},
		{
			callback: "a-b",
			resp:     []string{"a", "b"},
		},
		{
			callback: "a-b-c-c(kksksfd)-l",
			resp:     []string{"a", "b", "c", "c", "l"},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			symbolList, err := parseCallback(tCase.callback)
			if err != nil {
				fmt.Println(err)
				assert.Equal(t, tCase.haveError, true)
			} else {
				assert.Equal(t, len(tCase.resp), len(symbolList))
				for i := range tCase.resp {
					assert.Equal(t, tCase.resp[i], symbolList[i])
				}
			}
		})
	}
}

func Test_convertNumberToSymbol(t *testing.T) {
	for i, tCase := range []struct {
		num       int
		resp      string
		haveError bool
	}{
		{
			num:       -1,
			haveError: true,
		},
		{
			num:       26,
			haveError: true,
		},
		{
			num:  0,
			resp: "a",
		},
		{
			num:  10,
			resp: "k",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			val, err := convertNumberToSymbol(tCase.num)
			if err != nil {
				assert.Equal(t, tCase.haveError, true)
				return
			}
			assert.Equal(t, val, tCase.resp)

		})
	}
}

func Test_convertSymbolToNum(t *testing.T) {
	for i, tCase := range []struct {
		symbol    string
		resp      int
		haveError bool
	}{
		{
			symbol:    "kaka",
			haveError: true,
		},
		{
			symbol:    "!",
			haveError: true,
		},
		{
			symbol: "a",
			resp:   0,
		},
		{
			symbol: "z",
			resp:   25,
		},
		{
			symbol:    "",
			haveError: true,
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			val, err := convertSymbolToNum(tCase.symbol)
			if err != nil {
				assert.Equal(t, tCase.haveError, true)
				return
			}
			assert.Equal(t, val, tCase.resp)
		})
	}
}

func Test_NewNodesHandlerSimple(t *testing.T) {
	ctx := context.Background()
	template1 := []*Node{
		{
			HumanText: "button1",
			NextNodes: []*Node{
				{
					HumanText: "button3",
				},
			},
		},
		{
			HumanText: "button2",
			NextNodes: []*Node{
				{
					HumanText: "button4",
				},
				{
					HumanText: "button5",
				},
			},
		},
	}

	handler, err := NewNodesHandler(template1, "defaultMessage")
	if err != nil {
		t.Errorf("creating handler: %v", err)
		return
	}

	node, err := handler.GetNodeByCallback(ctx, 0, "a")
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	node2, err := handler.GetNodeByCallback(ctx, 0, "b")
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node.HumanText, "button1")
	assert.Equal(t, node2.HumanText, "button2")

	node3, err := handler.GetNodeByCallback(ctx, 0, "a-a")
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node3.HumanText, "button3")

	node4, err := handler.GetNodeByCallback(ctx, 0, "b-a")
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node4.HumanText, "button4")

	node5, err := handler.GetNodeByCallback(ctx, 0, "b-b")
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node5.HumanText, "button5")
}

func generateNodes(ctx context.Context, chatID int64) (NodesList, error) {
	return NodesList{
		{
			HumanText: "buttonInside1",
		},
		{
			HumanText: "buttonInside2",
		},
	}, nil
}

func Test_NewNodesHandlerNodesGenerating(t *testing.T) {
	ctx := context.Background()
	template1 := []*Node{
		{
			HumanText:          "button1",
			NextNodesGenerator: generateNodes,
		},
	}

	handler, err := NewNodesHandler(template1, "defaultMessage")
	if err != nil {
		t.Errorf("creating handler: %v", err)
		return
	}
	node, err := handler.GetNodeByCallback(ctx, 0, "a-a")
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node.HumanText, "buttonInside1")
	assert.Equal(t, node.Message, "defaultMessage")
}

func Test_GetCallbackBack(t *testing.T) {
	for i, tCase := range []struct {
		callback     string
		callbackBack string
		haveError    bool
	}{
		{
			callback:  "",
			haveError: true,
		},
		{
			callback:     "a",
			callbackBack: "",
		},
		{
			callback:     "a-b",
			callbackBack: "a",
		},
		{
			callback:     "a-b(10)-c",
			callbackBack: "a-b(10)",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var node Node
			node.callback = tCase.callback
			callbackBack, err := node.GetCallbackBack()
			if err != nil {
				assert.Equal(t, tCase.haveError, true)
			} else {
				assert.Equal(t, tCase.callbackBack, callbackBack)
			}
		})
	}
}
