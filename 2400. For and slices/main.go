package main

import (
	"fmt"
	"strings"
)

func FilterExpensiveOrders(orders []int, limit int) []int {
	tmp := []int{}
	for _, s := range orders {
		if limit < s {
			tmp = append(tmp, s)
		}
	}
	return tmp
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	nums := []int{1, 2, 3}

	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}

	fmt.Println(nums)

	langs := []string{"Go", "Rust", "Python"}

	for i, lang := range langs {
		fmt.Printf("langs [%d] = %s\n", i, lang)
	}

	for _, lang := range langs {
		fmt.Println(lang)
	}

	for i := range langs {
		fmt.Println(i)
	}

	for _, lang := range langs {
		lang = strings.ToUpper(lang)
	}

	fmt.Println(langs)

	for i := range langs {
		langs[i] = strings.ToUpper(langs[i])
	}

	fmt.Println(langs)

}
