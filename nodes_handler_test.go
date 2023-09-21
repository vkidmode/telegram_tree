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
		{
			name:    "11",
			element: "(10)",
			valid:   true,
		},
		{
			name:    "12",
			element: "!(10)",
			valid:   false,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			valid := checkCallbackElement(tCase.element)
			assert.Equal(t, tCase.valid, valid)
		})
	}
}

func Test_extractSymbolFromElem(t *testing.T) {
	for _, tCase := range []struct {
		element   string
		haveError bool
		symbol    string
		name      string
	}{
		{
			name:    "1",
			element: "z",
			symbol:  "z",
		},
		{
			name:      "2",
			element:   "z))",
			haveError: true,
		},
		{
			name:    "3",
			element: "a(ksksk)",
			symbol:  "a",
		},
		{
			name:      "4",
			element:   "aa(ksksk)",
			haveError: true,
		},
		{
			name:      "5",
			element:   "",
			haveError: true,
		},
		{
			name:      "6",
			element:   " ",
			haveError: true,
		},
		{
			name:    "6",
			element: "(s)",
			symbol:  "",
		},
		{
			element: fmt.Sprintf("%s(ksksk)", CallBackSkip),
			symbol:  CallBackSkip,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
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
		{
			name:     "9",
			callback: "a>(ss)>(ss)",
			resp:     []string{"a", "", ""},
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
			num:       301,
			haveError: true,
		},
		{
			num:  0,
			resp: "È",
		},
		{
			num:  10,
			resp: "Ò",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			val, err := ConvertNumberToSymbol(tCase.num)
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
			func(ctx context.Context, meta Meta) ([]Node, error) {
				return []Node{
					NewNode(NewTelegramOptions("", "button3", nil, false, false, false), nil, nil, nil),
				}, nil
			},
			nil,
		),

		NewNode(
			NewTelegramOptions("", "button2", nil, false, false, false),
			nil,
			func(ctx context.Context, meta Meta) ([]Node, error) {
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

func generateNodes(ctx context.Context, meta Meta) ([]Node, error) {
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
	for _, tCase := range []struct {
		callback     string
		callbackBack string
		haveError    bool
		name         string
	}{
		{
			name:      "1",
			callback:  "",
			haveError: true,
		},
		{
			name:         "2",
			callback:     "a",
			callbackBack: "",
		},
		{
			name:         "3",
			callback:     "a>b",
			callbackBack: "a",
		},
		{
			name:         "4",
			callback:     "a>b(10)>c",
			callbackBack: "a>b(10)",
		},
		{
			name:         "5",
			callback:     "a>c(10)>(4)",
			callbackBack: "a>c(10)",
		},
		{
			name:         "6",
			callback:     "a>c(10)>(4)>(10)",
			callbackBack: "a>c(10)>(4)",
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
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

func Test_GetNodeTreeOne(t *testing.T) {
	tree, err := NewNodesHandler([]Node{
		&node{
			telegramOptions: NewTelegramOptions("some", "I am root", nil, false, false, true),
			processor: func(ctx context.Context, meta Meta) ([]Node, error) {
				return []Node{
					&node{
						telegramOptions: NewTelegramOptions("some", "I am 2", nil, false, false, true),
						processor: func(ctx context.Context, meta Meta) ([]Node, error) {
							return []Node{
								&node{
									telegramOptions: NewTelegramOptions("some", "I am 3a", nil, false, false, true),
								},
								&node{
									telegramOptions: NewTelegramOptions("some", "I am 3b", nil, false, false, true),
								},
							}, nil
						},
					},
				}, nil
			},
		},
	}, "some msg")
	if err != nil {
		t.Error(err)
	}

	for _, tCase := range []struct {
		name      string
		callback  string
		humanText string
	}{
		{
			name:      "1",
			callback:  "a>a",
			humanText: "I am 2",
		},
		{
			name:      "2",
			callback:  "a>a>a",
			humanText: "I am 3a",
		},
		{
			name:      "3",
			callback:  "a>a>b",
			humanText: "I am 3b",
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.Background()
			nodeData, err := tree.GetNode(ctx, newMeta(tCase.callback))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tCase.humanText, nodeData.GetTelegramOptions().GetHumanText())
		})
	}
}

func Test_GetNodeTree2(t *testing.T) {
	tree, err := genNewTree()
	if err != nil {
		t.Error(err)
	}

	for _, tCase := range []struct {
		name      string
		callback  string
		humanText string
	}{
		{
			name:      "1",
			callback:  "a>a>a(.=4)",
			humanText: "testReg1",
		},
		{
			name:      "2",
			callback:  "a>a>a(.=4)>b",
			humanText: "testSubReg2",
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.Background()
			nodeData, err := tree.GetNode(ctx, newMeta(tCase.callback))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tCase.humanText, nodeData.GetTelegramOptions().GetHumanText())
		})
	}
}

func generateCountryCityRoot() Node {
	return NewNode(
		NewTelegramOptions("", "NONE", nil, true, false, true),
		nil,
		generateCountrySearchNodes,
		nil,
	)
}

func generateCountrySearchNodes(_ context.Context, info Meta) ([]Node, error) {
	return []Node{
		NewNode(
			NewTelegramOptions("", "select country/city from list", nil, false, false, true),
			nil,
			generateRegions,
			nil,
		),
		NewNode(
			NewTelegramOptions("", "Use my geoip", nil, true, false, true),
			nil,
			nil,
			nil,
		),
	}, nil
}

type regions struct {
	id   int64
	name string
}

func generateRegions(_ context.Context, _ Meta) ([]Node, error) {
	reg := []regions{
		{
			id:   1,
			name: "testReg1",
		},
		{
			id:   2,
			name: "testReg2",
		},
	}

	out := make([]Node, len(reg))

	for i := range reg {
		out[i] = NewNode(
			NewTelegramOptions("Select your region", reg[i].name, nil, false, false, true),
			NewPayload(".", strconv.Itoa(int(reg[i].id))),
			generateSubregions,
			nil,
		)
	}
	return out, nil
}

func generateSubregions(ctx context.Context, info Meta) ([]Node, error) {
	payloadData, err := ExtractPayload(info.GetCallback())
	if err != nil {
		return nil, err
	}

	_, ok := payloadData["."]
	if !ok {
		return nil, err
	}

	subReg := []regions{
		{
			id:   1,
			name: "testSubReg1",
		},
		{
			id:   2,
			name: "testSubReg2",
		},
	}

	out := make([]Node, len(subReg))

	for i := range subReg {
		out[i] = NewNode(
			NewTelegramOptions("Select your subregion", subReg[i].name, nil, false, false, true),
			nil,
			nil,
			nil,
		)
	}
	return out, nil
}

func genNewTree() (treeHandler *NodesHandler, err error) {
	treeHandler, err = NewNodesHandler(generateRootNodes(), "Выбери:")
	if err != nil {
		return nil, err
	}
	if treeHandler == nil {
		return nil, err
	}
	return treeHandler, nil
}

func generateRootNodes() []Node {
	return []Node{
		generateCountryCityRoot(),
	}
}
