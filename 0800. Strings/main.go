package main

import "fmt"
import "strconv"

func BuildGreeting(name string, count int) string {
	s1 := "Hello, "
	s2 := "! You have "
	s3 := " new messages."
	return s1 + name + s2 + strconv.Itoa(count) + s3
}

// func BuildGreeting(name string, count int) string {
// 	return fmt.Sprintf("Hello, %s! You have %d new messages.", name, count)
// }

func main() {
	// text := "Hello"
	// text += " World"
	// fmt.Println(text) // Hello World

	// import "unicode/utf8"
	// utf8.RuneCountInString("хекслет") // 6

	runes := []rune("Привет")

	// код точки Unicode
	fmt.Println(runes[0])        // => 1055
	fmt.Printf("%c\n", runes[0]) // => П

	fmt.Println(BuildGreeting("Dmitry", 5))
}
