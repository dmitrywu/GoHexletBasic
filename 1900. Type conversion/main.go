package main

import "fmt"

func BuildProfile(name string, age int, rating float64) string {
	return fmt.Sprintf("Name: %s, Age: %d, Rating: %.1f", name, age, rating)
}

func main() {
	fmt.Println(BuildProfile("Alice", 30, 4.73))
	// => "Name: Alice, Age: 30, Rating: 4.7"

	fmt.Println(BuildProfile("Bob", 25, 5))
	// => "Name: Bob, Age: 25, Rating: 5.0"
}
