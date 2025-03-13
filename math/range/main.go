package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rang(nums []int) int {
	max := nums[0]
	min := nums[0]
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max - min
}
func main() {
	var t [5]int
	rand.Seed(time.Now().UnixNano())
	for i := range t {
		t[i] = rand.Intn(100)
	}
	fmt.Println(t)
	fmt.Println(rang(t[:]))
}
