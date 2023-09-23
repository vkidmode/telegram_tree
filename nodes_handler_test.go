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
			element: "È",
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
			element: "È(dasgljsdfgn2!)",
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
			element: "È(.=2023-03-12)",
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
			element: "Ñ",
			symbol:  "Ñ",
		},
		{
			name:      "2",
			element:   "Ñ))",
			haveError: true,
		},
		{
			name:    "3",
			element: "È(ksksk)",
			symbol:  "È",
		},
		{
			name:      "4",
			element:   "ÈÈ(ksksk)",
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
			element: "(È)",
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
			callback: "È",
			resp:     []string{"È"},
		},
		{
			name:      "3",
			callback:  "È>>>>",
			haveError: true,
		},
		{
			name:      "4",
			callback:  "È>>É",
			haveError: true,
		},
		{
			name:     "5",
			callback: "È>É",
			resp:     []string{"È", "É"},
		},
		{
			name:     "6",
			callback: "È>É>Ï>Ï(kksksfd)>Ñ",
			resp:     []string{"È", "É", "Ï", "Ï", "Ñ"},
		},
		{
			name:     "7",
			callback: fmt.Sprintf("È>É>%s>Ï(kksksfd)>Ñ", CallBackSkip),
			resp:     []string{"È", "É", CallBackSkip, "Ï", "Ñ"},
		},
		{
			name:     "8",
			callback: "È>È>É>@>È>È(.=2023-03-12)",
			resp:     []string{"È", "È", "É", CallBackSkip, "È", "È"},
		},
		{
			name:     "9",
			callback: "È>(ss)>(ss)",
			resp:     []string{"È", "", ""},
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
	for _, tCase := range []struct {
		symbol    string
		resp      int
		haveError bool
		name      string
	}{
		{
			name:      "1",
			symbol:    "kaka",
			haveError: true,
		},
		{
			name:      "2",
			symbol:    "!",
			haveError: true,
		},
		{
			name:   "3",
			symbol: "È",
			resp:   0,
		},
		{
			name:   "4",
			symbol: "á",
			resp:   25,
		},
		{
			name:      "5",
			symbol:    "",
			haveError: true,
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
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
			WithTg(NewTelegram(WithTabTxt("button1"))),
			WithProc(
				func(ctx context.Context, meta Meta) ([]Node, error) {
					return []Node{
						NewNode(WithTg(NewTelegram(WithTabTxt("button3")))),
					}, nil
				},
			),
		),

		NewNode(
			WithTg(NewTelegram(WithTabTxt("button2"))),
			WithProc(
				func(ctx context.Context, meta Meta) ([]Node, error) {
					return []Node{
						NewNode(WithTg(NewTelegram(WithTabTxt("button4")))),
						NewNode(WithTg(NewTelegram(WithTabTxt("button5")))),
					}, nil
				},
			),
		),
	}

	handler, err := NewNodesHandler(template1, "defaultMessage")
	if err != nil {
		t.Errorf("creating handler: %v", err)
		return
	}

	node1, err := handler.GetNode(ctx, newMeta("È"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node1.GetTelegram().GetTabTxt(), "button1")

	node2, err := handler.GetNode(ctx, newMeta("É"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node2.GetTelegram().GetTabTxt(), "button2")

	node3, err := handler.GetNode(ctx, newMeta("È>È"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node3.GetTelegram().GetTabTxt(), "button3")

	node4, err := handler.GetNode(ctx, newMeta("É>È"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node4.GetTelegram().GetTabTxt(), "button4")

	node5, err := handler.GetNode(ctx, newMeta("É>É"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, node5.GetTelegram().GetTabTxt(), "button5")
}

func generateNodes(ctx context.Context, meta Meta) ([]Node, error) {
	return []Node{
		NewNode(WithTg(NewTelegram(WithTabTxt("buttonInside1")))),
		NewNode(WithTg(NewTelegram(WithTabTxt("buttonInside2")))),
	}, nil
}

func Test_NewNodesHandlerNodesGenerating(t *testing.T) {
	ctx := context.Background()

	template1 := nodesList{
		NewNode(WithTg(NewTelegram(WithTabTxt("button1"))), WithProc(generateNodes)),
	}

	handler, err := NewNodesHandler(template1, "defaultMessage")
	if err != nil {
		t.Errorf("creating handler: %v", err)
		return
	}
	nodeItem, err := handler.GetNode(ctx, newMeta("È>È"))
	if err != nil {
		t.Errorf("getting node by callback: %v", err)
		return
	}
	assert.Equal(t, nodeItem.GetTelegram().GetTabTxt(), "buttonInside1")
	assert.Equal(t, nodeItem.GetTelegram().GetMessage(), "defaultMessage")
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
			callback:     "È",
			callbackBack: "",
		},
		{
			name:         "3",
			callback:     "È>É",
			callbackBack: "È",
		},
		{
			name:         "4",
			callback:     "È>É(10)>Ê",
			callbackBack: "È>É(10)",
		},
		{
			name:         "5",
			callback:     "È>Ê(10)>(4)",
			callbackBack: "È>Ê(10)",
		},
		{
			name:         "6",
			callback:     "È>Ê(10)>(4)>(10)",
			callbackBack: "È>Ê(10)>(4)",
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
			telegram: NewTelegram(WithMessage("some"), WithTabTxt("I am root")),
			processor: func(ctx context.Context, meta Meta) ([]Node, error) {
				return []Node{
					&node{
						telegram: NewTelegram(WithMessage("some"), WithTabTxt("I am 2")),
						processor: func(ctx context.Context, meta Meta) ([]Node, error) {
							return []Node{
								&node{
									telegram: NewTelegram(WithMessage("some"), WithTabTxt("I am 3a")),
								},
								&node{
									telegram: NewTelegram(WithMessage("some"), WithTabTxt("I am 3b")),
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
			callback:  "È>È",
			humanText: "I am 2",
		},
		{
			name:      "2",
			callback:  "È>È>È",
			humanText: "I am 3a",
		},
		{
			name:      "3",
			callback:  "È>È>É",
			humanText: "I am 3b",
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.Background()
			nodeData, err := tree.GetNode(ctx, newMeta(tCase.callback))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tCase.humanText, nodeData.GetTelegram().GetTabTxt())
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
			callback:  "È>È>È(.=4)",
			humanText: "testReg1",
		},
		{
			name:      "2",
			callback:  "È>È>È(.=4)>É",
			humanText: "testSubReg2",
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			ctx := context.Background()
			nodeData, err := tree.GetNode(ctx, newMeta(tCase.callback))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, tCase.humanText, nodeData.GetTelegram().GetTabTxt())
		})
	}
}

func generateCountryCityRoot() Node {
	return NewNode(WithTg(NewTelegram(WithTabTxt("NONE"))), WithProc(generateCountrySearchNodes))
}

func generateCountrySearchNodes(_ context.Context, info Meta) ([]Node, error) {
	return []Node{
		NewNode(
			WithTg(NewTelegram(WithTabTxt("choose country/city from list"))),
			WithProc(generateRegions),
		),
		NewNode(
			WithTg(NewTelegram(WithTabTxt("Use my geoip"))),
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
			WithTg(NewTelegram(WithMessage("Select your region"), WithTabTxt(reg[i].name))),
			WithPayload(NewPayload(".", strconv.Itoa(int(reg[i].id)))),
			WithProc(generateSubregions),
		)
	}
	return out, nil
}

func generateSubregions(_ context.Context, info Meta) ([]Node, error) {
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
		out[i] = NewNode(WithTg(NewTelegram(WithMessage("Select your subregion"), WithTabTxt(subReg[i].name))))
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
