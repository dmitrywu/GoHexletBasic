package main

import "fmt"

func RepeatWord(word string, times int) string {
	if times == 1 {
		return word
	}
	var strWord string
	for range times {
		strWord += word
	}
	return strWord
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	text := "go"

	for i, r := range text {
		fmt.Printf("Индекс: %d, Символ: %c\n", i, r)
	}

	word := "го"

	for i, r := range word {
		fmt.Println(i, string(r))
	}

	for _, r := range "hexlet" {
		fmt.Printf("%c\n", r)
	}

	i := 0

	nums := make([]int, 0, 10)

	for {
		i++
		if i < 2 {
			i++
		}
		if i == 9 {
			break
		}
		nums = append(nums, i)
	}
	fmt.Println(nums)

	fmt.Println(RepeatWord("hi", 5))
}
