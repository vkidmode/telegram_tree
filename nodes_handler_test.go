package telegram_tree

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkCallbackElement(t *testing.T) {
	for _, tCase := range []struct {
		name    string
		element string
		valid   bool
	}{
		{
			name:    "1",
			element: "kaka",
			valid:   false,
		},
		{
			name:    "2",
			element: "a",
			valid:   true,
		},
		{
			name:    "3",
			element: "",
			valid:   false,
		},
		{
			name:    "4",
			element: "A",
			valid:   false,
		},
		{
			name:    "5",
			element: "z(dasgljsdfgn2!)",
			valid:   true,
		},
		{
			name:    "6",
			element: "z(aaaa)z",
			valid:   false,
		},
		{
			name:    "7",
			element: "z()",
			valid:   false,
		},
		{
			name:    "8",
			element: "()",
			valid:   false,
		},
		{
			name:    "9",
			element: fmt.Sprintf("%s(kaka)", CallBackSkip),
			valid:   true,
		},
		{
			name:    "10",
			element: "a(.=2023-03-12)",
			valid:   true,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
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
		{
			element: fmt.Sprintf("%s(ksksk)", CallBackSkip),
			symbol:  CallBackSkip,
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
	for _, tCase := range []struct {
		name      string
		callback  string
		haveError bool
		resp      []string
	}{
		{
			name:      "1",
			callback:  "",
			haveError: true,
		},
		{
			name:     "2",
			callback: "a",
			resp:     []string{"a"},
		},
		{
			name:      "3",
			callback:  "a>>>>",
			haveError: true,
		},
		{
			name:      "4",
			callback:  "a>>b",
			haveError: true,
		},
		{
			name:     "5",
			callback: "a>b",
			resp:     []string{"a", "b"},
		},
		{
			name:     "6",
			callback: "a>b>c>c(kksksfd)>l",
			resp:     []string{"a", "b", "c", "c", "l"},
		},
		{
			name:     "7",
			callback: fmt.Sprintf("a>b>%s>c(kksksfd)>l", CallBackSkip),
			resp:     []string{"a", "b", CallBackSkip, "c", "l"},
		},
		{
			name:     "8",
			callback: "a>a>b>@>a>a(.=2023-03-12)",
			resp:     []string{"a", "a", "b", CallBackSkip, "a", "a"},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			symbolList, err := parseCallback(tCase.callback)
			if err != nil {
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
	template1 := nodesList{
		NewNode(
			NewTelegramOptions("", "button1", nil, false, false, false),
			nil,
			func(ctx context.Context, meta meta) ([]Node, error) {
				return []Node{
					NewNode(NewTelegramOptions("", "button3", nil, false, false, false), nil, nil, nil),
				}, nil
			},
			nil,
		),

		NewNode(
			NewTelegramOptions("", "button2", nil, false, false, false),
			nil,
			func(ctx context.Context, meta meta) ([]Node, error) {
				return []Node{
					NewNode(NewTelegramOptions("", "button4", nil, false, false, false), nil, nil, nil),
					NewNode(NewTelegramOptions("", "button5", nil, false, false, false), nil, nil, nil),
				}, nil
			},
			nil,
		),
	}

	handler, err := NewNodesHandler(template1, "defaultMessage")
	if err != nil {
		t.Errorf("creating handler: %v", err)
		return
	}

	nodeItem, err := handler.GetNode(ctx, newMeta("a"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	node2, err := handler.GetNode(ctx, newMeta("b"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, nodeItem.GetTelegramOptions().GetHumanText(), "button1")
	assert.Equal(t, node2.GetTelegramOptions().GetHumanText(), "button2")

	node3, err := handler.GetNode(ctx, newMeta("a>a"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node3.GetTelegramOptions().GetHumanText(), "button3")

	node4, err := handler.GetNode(ctx, newMeta("b>a"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node4.GetTelegramOptions().GetHumanText(), "button4")

	node5, err := handler.GetNode(ctx, newMeta("b>b"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node5.GetTelegramOptions().GetHumanText(), "button5")
}

func generateNodes(ctx context.Context, meta meta) ([]Node, error) {
	return []Node{
		NewNode(NewTelegramOptions("", "buttonInside1", nil, false, false, false), nil, nil, nil),
		NewNode(NewTelegramOptions("", "buttonInside2", nil, false, false, false), nil, nil, nil),
	}, nil
}

func Test_NewNodesHandlerNodesGenerating(t *testing.T) {
	ctx := context.Background()

	template1 := nodesList{
		NewNode(NewTelegramOptions("", "button1", nil, false, false, false), nil, generateNodes, nil),
	}

	handler, err := NewNodesHandler(template1, "defaultMessage")
	if err != nil {
		t.Errorf("creating handler: %v", err)
		return
	}
	nodeItem, err := handler.GetNode(ctx, newMeta("a>a"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, nodeItem.GetTelegramOptions().GetHumanText(), "buttonInside1")
	assert.Equal(t, nodeItem.GetTelegramOptions().GetMessage(), "defaultMessage")
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
			callback:     "a>b",
			callbackBack: "a",
		},
		{
			callback:     "a>b(10)>c",
			callbackBack: "a>b(10)",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var nodeItem node
			nodeItem.callback = tCase.callback
			callbackBack, err := nodeItem.GetCallbackBack()
			if err != nil {
				assert.Equal(t, tCase.haveError, true)
			} else {
				assert.Equal(t, tCase.callbackBack, callbackBack)
			}
		})
	}
}
