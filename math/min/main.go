package main

import "fmt"

func min(nums []int) int {
	m:=nums[0]
	for i:=range nums{
		if nums[i]<m{
			m=nums[i]
		}
	}
	return m
}

func main() {
	fmt.Println(min([]int{3,4,5,1,7,8}))
}
