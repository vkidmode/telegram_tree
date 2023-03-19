package telegram_tree

import (
	"fmt"
	"regexp"
	"strings"
)

type callbackSymbolsList []string

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
	return regexp.MatchString(`^[a-z+@](\(.+\))?$`, element)
}

func incrementCallback(callback string, payload Payload, number int) (string, error) {
	symbol, err := convertNumberToSymbol(number)
	if err != nil {
		return "", err
	}
	if callback == "" {
		return symbol, nil
	}
	if _, err = parseCallback(callback); err != nil {
		return "", err
	}
	resp := fmt.Sprintf("%s%s%s", callback, callbackDivider, symbol)
	if payload.GetValue() != "" && payload.GetKey() != "" {
		resp = fmt.Sprintf("%s(%s=%s)", resp, payload.GetKey(), payload.GetValue())
	}
	return resp, nil
}
