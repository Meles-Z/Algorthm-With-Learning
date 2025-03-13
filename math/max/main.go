package main

import "fmt"

func maxValue(nums []int) int {
	fmt.Println(nums)
	max := nums[0]
	for i, _ := range nums[1:] {
		if nums[i+1] > max {
			max = nums[i+1]
		}
		fmt.Println(nums[i+1])
	}
	return max
}
func main() {
	fmt.Println(maxValue([]int{2, 3, 10, 4, 4, 5, 6}))
}
