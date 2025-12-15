package main

import (
	"fmt"
	"slices"
)

func AreOrderHistoriesEqual(history1, history2 [][]string) bool {
	if len(history1) != len(history2) {
		return false
	}

	if history1 == nil || history2 == nil {
		return false
	}

	for i := range history1 {
		if len(history1[i]) != len(history2[i]) {
			return false
		}

		for j := range history1[i] {
			if history1[i][j] != history2[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	fmt.Println(slices.Equal(a, b))

	c := []int{1, 2, 3}
	d := c
	fmt.Println(&c[0] == &d[0])

	var s []int
	fmt.Println(s == nil)

	empty := []int{}
	fmt.Println(empty == nil)

	h1 := [][]string{
		{"milk", "bread"},
		{"apple", "banana"}}

	h2 := [][]string{
		{"milk", "bread", "butter"},
		{"apple", "banana"}}

	fmt.Println("----")
	fmt.Println(AreOrderHistoriesEqual(h1, h2))
}
