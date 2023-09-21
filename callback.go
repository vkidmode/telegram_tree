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
	valid := checkCallbackElement(in)
	if !valid {
		return "", fmt.Errorf("element is not valid")
	}
	if in[0] != '(' {
		return string(in[0]), nil
	}
	return "", nil
}

func checkCallbackElement(element string) bool {
	if element == CallBackSkip {
		return true
	}
	if len(element) == 0 {
		return false
	}
	if element[0] != '(' {
		if string(element[0]) != CallBackSkip {
			if _, err := convertSymbolToNum(string(element[0])); err != nil {
				return false
			}
		}
		if len(element) == 1 {
			return true
		}
		return checkBraces(element[1:])
	}
	return checkBraces(element)
}

func checkBraces(in string) bool {
	if len(in) < 3 {
		return false
	}
	if in[0] == '(' && in[len(in)-1] == ')' {
		return true
	}
	return false
}

func incrementCallback(callback string, payload Payload, number int) (string, error) {
	symbol, err := ConvertNumberToSymbol(number)
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
	if payload != nil {
		if payload.GetValue() != "" && payload.GetKey() != "" {
			resp = fmt.Sprintf("%s(%s=%s)", resp, payload.GetKey(), payload.GetValue())
		}
	}
	return resp, nil
}

func extractPayloadFromCallback(callback string) (map[string]string, error) {
	if _, err := parseCallback(callback); err != nil {
		return nil, err
	}

	var resp = make(map[string]string)
	r := regexp.MustCompile(`\(.*?\)`)
	matches := r.FindAllString(callback, -1)
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
