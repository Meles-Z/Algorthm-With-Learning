package main

import "fmt"

func squareOfNumber(nums []int) []int {
	var result []int
	// traverse on each number and check every node
	for i := range nums {
		result = append(result, nums[i]*nums[i])
	}
	return result
}

func squareRoot(n int) int {
	var result int
	if n < 0 {
		return -1
	}
	for i := 1; i*i <= n; i++ {

		result = i
	}
	return result
}
func main() {
	fmt.Println("Hello, world!")
	fmt.Println(squareOfNumber([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(squareRoot(16))
}
