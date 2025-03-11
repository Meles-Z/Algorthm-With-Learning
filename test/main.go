package main

import (
	"fmt"
	"math"
)

func squareOfNumber(nums []int) []int {
	var result []int
	// traverse on each number and check every node
	for i := range nums {
		result = append(result, nums[i]*nums[i])
	}
	return result
}

func squareRoot(n float64) float64 {
	if n < 0 {
		return -1
	}
	return math.Sqrt(float64(n))
}
func main() {
	fmt.Println("Hello, world!")
	fmt.Println(squareOfNumber([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(squareRoot(8))
}
