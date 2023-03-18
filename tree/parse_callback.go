package tree

import (
	"fmt"
	"regexp"
	"strings"
)

const callbackDivider = "-"

type callbackSymbolsList []string

func (c callbackSymbolsList) toNumbers() ([]int, error) {
	resp := make([]int, len(c))
	for i := range c {
		num, err := convertSymbolToNum(c[i])
		if err != nil {
			return nil, err
		}
		resp[i] = num
	}
}

func parseCallback(callback string) (callbackSymbolsList, error) {
	if callback == "" {
		return nil, fmt.Errorf("empty callback")
	}
	if strings.Contains(callback, " ") {
		return nil, fmt.Errorf("invalid callback")
	}
	callbackElements := strings.Split(callback, callbackDivider)
	var resp = make([]string, 0, len(callbackElements))
	for i := range callbackElements {
		symbol, err := extractSymbolFromElem(callbackElements[i])
		if err != nil {
			return nil, err
		}
		resp = append(resp, symbol)
	}
	return resp, nil
}

func extractSymbolFromElem(in string) (string, error) {
	valid, err := checkCallbackElement(in)
	if err != nil {
		return "", err
	}
	if !valid {
		return "", fmt.Errorf("element is not valid")
	}
	return string(in[0]), nil
}

func checkCallbackElement(element string) (bool, error) {
	return regexp.MatchString(`^[a-z](\(.+\))?$`, element)
}
