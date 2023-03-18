package tree

import (
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
