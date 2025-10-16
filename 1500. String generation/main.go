package main

import (
	"fmt"
	//"strings"
	//	"unicode"
)

func FilterString(text string, symb rune) string {
	if text == "" {
		return text
	}

	// var sb strings.Builder

	// var r1 byte = byte(unicode.ToLower(symb))
	// var r2 byte = byte(unicode.ToUpper(symb))

	for _, i := range text {
		fmt.Print(i)
		fmt.Println(text)

	}

	return ""

}

func main() {

	fmt.Println(FilterString("If I look forward I win", 'i'))

}
