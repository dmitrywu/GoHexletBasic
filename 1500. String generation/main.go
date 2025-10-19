package main

import (
	"fmt"
	"strings"
	"unicode"
)

func FilterString(text string, symb rune) string {
	if text == "" {
		return text
	}

	var sb strings.Builder

	s := unicode.ToLower(symb)

	for _, r := range text {
		if unicode.ToLower(r) != s {
			sb.WriteRune(r)
		}
	}

	return sb.String()

}

func main() {

	fmt.Println(FilterString("", 'i'))
	fmt.Println(FilterString("If I look forward I win", 'i'))

}
