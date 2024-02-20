package propertybased

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
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

func ConvertToRoman(n int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			result.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(r string) int {
	var arabic int = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(r, numeral.Symbol) {
			arabic += numeral.Value
			r = strings.TrimPrefix(r, numeral.Symbol)
		}
	}

	return arabic
}
