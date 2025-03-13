package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func media(nums []int) float64 {
	sort.Ints(nums)
	fmt.Println("generated numbers:", nums)
	mid := len(nums) / 2

	if len(nums)%2 != 0 {
		return float64(nums[mid])
	}
	return (float64((nums[mid-1])) + float64(nums[mid])) / 2
}
func main() {
	var v [5]int
	rand.Seed(time.Now().UnixNano())
	for i := range v {
		v[i] = rand.Intn(100)
	}
	fmt.Println(media(v[:]))
}
