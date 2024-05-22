package converter

import (
	"strings"
	"fmt"
)

var romanNumerals = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRoman(num int) (string, error) {
	if num < 1 || num > 3999 {
		return "", fmt.Errorf("number %d out of range (1-3999)", num)
	}
	var result strings.Builder
	for _, roman := range romanNumerals {
		count := num / roman.Value
		num %= roman.Value
		result.WriteString(strings.Repeat(roman.Symbol, count))
	}
	return result.String(), nil
}
