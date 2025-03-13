package main

import "fmt"

func mode(nums []int) []int {
	freq := make(map[int]int)
	maxFreq := 0
	modes := []int{}
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
		}
	}
	for i, v := range freq {
		if v == maxFreq {
			modes = append(modes, i)
		}
	}
	return modes
}
func main() {
	fmt.Println(mode([]int{3, 4, 5, 5, 5, 6, 6, 6, 8, 8,8}))
}
