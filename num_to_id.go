package telegram_tree

import (
	"fmt"
)

func convertNumberToSymbol(in int) (string, error) {
	if in < 0 || in > 25 {
		return "", fmt.Errorf("unsupported number")
	}
	return string(rune(in + 97)), nil
}

func convertSymbolToNum(in string) (int, error) {
	if len(in) != 1 {
		return 0, fmt.Errorf("unsopported symbol")
	}
	num := int(in[0])
	if num < 97 || num > 122 {
		return 0, fmt.Errorf("unsopported symbol")
	}
	return num - 97, nil
}
