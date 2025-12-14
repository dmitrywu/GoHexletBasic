package main

import (
	"fmt"
	"slices"
)

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

}
