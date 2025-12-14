package main

func GetWorkHours(schedule []int) []int {
	if len(schedule) < 5 {
		return []int{}
	}
	sum := make([]int, 5)
	// for _, v := range(schedule[:5]){
	// 	sum =+v
	// }
	copy(sum, schedule)
	return sum
}

func AddDiscount(prices []float64, discount float64) []float64 {
	length := len(prices)
	if length < 1 {
		return prices
	}
	prices[length-1] = prices[length-1] * (1 - (discount / 100))
	return prices
}

func main() {
	var num []int
	numbers := []int{}
	numbers1 := []int{10, 10, 30}
	nums := make([]int, 5)
	buf := make([]int, 0, 1000)
	nums = append(nums, 3)
	nums = append(nums, 2, 1, 0)

}
