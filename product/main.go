package main

import "fmt"

func maxProduct(nums []int) int {
	result := 0
	if len(nums) == 1 {
		result = nums[0]
	}
	for i, _ := range nums {
		if nums[i] > result {
			result = nums[i]
		}
		fmt.Println(result)
		if nums[i]*nums[i+1] > result {
			result = nums[i] * nums[i+1]
		}
	}
	return result
}
func main() {
	fmt.Println(maxProduct([]int{0, 2}))
}
