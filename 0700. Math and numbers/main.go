package main

import "fmt"

func CalculateProgress(done int, total int) float64 {
	return float64(done) / float64(total)
}

func main() {
	// number type in Go:
	// uint, uint8, uint16, uint32, uint64
	// int, int8, int16, int32, int64
	// float32, float64
	// complex64, complex128

	count := 10
	count /= 4
	fmt.Println(count)

	x := int64(5.0)

	fmt.Println(x)

	fmt.Println(CalculateProgress(0, 10))
	fmt.Println(CalculateProgress(1, 4))
	fmt.Println(CalculateProgress(2, 5))
	fmt.Println(CalculateProgress(3, 4))
	fmt.Println(CalculateProgress(5, 5))
}
