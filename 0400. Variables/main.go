package main

import "fmt"

func main() {
	var name string
	name = "Vasya"
	println(name)

	// Default values
	// string - ""
	// int - 0
	// bool - false

	var nameTwo = "Hexlet"
	var count = 10
	var active = true

	fmt.Println(nameTwo, count, active)

	txt := "Hex"
	text, course := "Learning", "Go"

	fmt.Println(txt, text, course)

}
