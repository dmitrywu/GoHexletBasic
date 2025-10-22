package main

import (
	"errors"
	"fmt"
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

// Реализуйте функцию GetFileExtension(filename string) (string, error), которая возвращает расширение файла
// (текст после последней точки). Если в имени файла нет точки, функция должна вернуть ошибку.
// файл "LICENSE" не имеет расширения
func GetFileExtension(filename string) (string, error) {

	findDot := -1

	for i := 0; i < len(filename); i++ {
		if filename[i] == '.' {
			findDot = i
		}
	}

	if findDot == -1 || findDot == len(filename)-1 {
		s := fmt.Sprintf("файл %s не имеет расширения", filename)
		return "", errors.New(s)
	}

	return filename[findDot+1:], nil
}

func main() {

	orders := []int{120, 35, 70, 400, 15, 220, 90}
	fmt.Println(FilterExpensiveOrders(orders, 100))            // [120 400 220]
	fmt.Println(FilterExpensiveOrders([]int{}, 100))           // []
	fmt.Println(FilterExpensiveOrders([]int{10, 20, 30}, 100)) // []

	ext, err := GetFileExtension("photo.jpeg")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Расширение файла:", ext)
	}
	// => Расширение файла: jpeg

	ext, err = GetFileExtension("backup.zip")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Расширение файла:", ext)
	}
	// => Расширение файла: zip

	ext, err = GetFileExtension("LICENSE")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Расширение файла:", ext)
	}
	// => Ошибка: файл "LICENSE" не имеет расширения

}
