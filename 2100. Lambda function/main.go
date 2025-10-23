package main

import "fmt"

// MakeGreeting принимает приветствие и возвращает новую функцию,
// которая добавляет имя к этому приветствию.
func MakeGreeting(greeting string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
}

func main() {
	var m int
	fmt.Println(m)
	greeter := MakeGreeting("Hello")
	fmt.Println(greeter("Hexlet")) // Hello, Hexlet!

	hiGreeter := MakeGreeting("Hi")
	fmt.Println(hiGreeter("Go")) // Hi, Go!
}
