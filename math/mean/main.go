package main

import "fmt"

func mean(nums []int) float64 {
	// create variables that can store all sum
	var sum = 0
	for i := range nums {
		sum += nums[i]
	}

	return float64(sum) / float64(len(nums))

}
func main() {
	fmt.Println("Welcome back!")
	fmt.Println(mean([]int{1,2,3,4}))
}
