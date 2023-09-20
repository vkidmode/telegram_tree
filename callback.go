package telegram_tree

import (
	"fmt"
	"regexp"
	"strings"
)

type callbackSymbolsList []string

type Callback string

func (c Callback) parseCallback() (callbackSymbolsList, error) {
	if c == "" {
		return nil, fmt.Errorf("empty callback")
	}
	callbackElements := strings.Split(c.String(), callbackDivider)
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

func (c Callback) String() string {
	return string(c)
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

func incrementCallback(callback Callback, payload Payload, number int) (Callback, error) {
	symbol, err := convertNumberToSymbol(number)
	if err != nil {
		return "", err
	}
	if callback == "" {
		return Callback(symbol), nil
	}

	if _, err = callback.parseCallback(); err != nil {
		return "", err
	}
	resp := Callback(fmt.Sprintf("%s%s%s", callback, callbackDivider, symbol))
	if payload != nil {
		if payload.GetValue() != "" && payload.GetKey() != "" {
			resp = Callback(fmt.Sprintf("%s(%s=%s)", resp, payload.GetKey(), payload.GetValue()))
		}
	}
	return resp, nil
}

func extractPayloadFromCallback(callback Callback) (map[string]string, error) {
	if _, err := callback.parseCallback(); err != nil {
		return nil, err
	}

	var resp = make(map[string]string)
	r := regexp.MustCompile(`\(.*?\)`)
	matches := r.FindAllString(callback.String(), -1)
	for i := range matches {
		matches[i] = strings.ReplaceAll(matches[i], ")", "")
		matches[i] = strings.ReplaceAll(matches[i], "(", "")

		substrings := strings.Split(matches[i], "=")
		if len(substrings) != 2 {
			return nil, fmt.Errorf("invalid payload")
		}
		resp[substrings[0]] = substrings[1]
	}
	return resp, nil
}
