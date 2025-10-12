package main

import (
	"fmt"
	"strings"
)

func GetHiddenCard(text string, count int) string {
	return strings.Repeat("*", count) + text[12:]
}

func main() {

	fmt.Println(GetHiddenCard("1234567812345678", 4))

	strings.ToLower("HeLLo") // "hello"
	strings.ToUpper("HeLLo") // "HELLO"

	strings.HasPrefix("golang", "go")      // true
	strings.HasSuffix("version.go", ".go") // true

	strings.Contains("hello world", "world") // true
	strings.Index("hello world", "lo")       // 3
	strings.Index("hello world", "go")       // -1

	strings.Replace("foo bar foo", "foo", "baz", 1)  // "baz bar foo"
	strings.Replace("foo bar foo", "foo", "baz", -1) // "baz bar baz

	strings.Repeat("na", 3) // "nananana"
	fmt.Println(strings.Repeat("na", 3))

	strings.TrimSpace("  hello\n") // "hello"

	text := "golang"
	// символы с позиции 0 до 2, не включая 2
	fmt.Println(text[0:2]) // => "go"
	// от начала до позиции 4
	fmt.Println(text[:4]) // => "gola"
	// от позиции 3 до конца строки
	fmt.Println(text[3:]) // => "ang"
	// вся строка
	fmt.Println(text[:]) // => "golang"

}
