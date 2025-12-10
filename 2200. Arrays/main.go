package main

import (
	"fmt"
)

func reverse(text [5]rune) [5]rune {
	var result [5]rune
	for i := 0; i < len(text); i++ {
		result[i] = text[len(text)-1-i]
	}
	return result
}

func SafeWrite(nums [5]int, i, val int) [5]int {
	if i >= 0 && i < len(nums) {
		nums[i] = val
		return nums
	}
	return nums
}

func main() {
	// Массив - фиксированное количество элементов ОДНОГО типа
	var nums [5]int
	nms := [5]int{}             // в этом случае скобки обязательны
	bms := [5]int{1, 2}         // оставшиеся инициализируются 0-ми
	xms := [...]int{10, 20, 30} // при использовании "..." длина посчитается автоматически

	fmt.Println(nums, nms, bms, xms, len(xms))

	// в функции массивы передаются по значению

	text := [5]rune{'п', 'р', 'и', 'в', 'е'}
	reversed := reverse(text)
	fmt.Println(string(reversed[:]))

}
