package mymath

import (
	"fmt"
	"sort"
)

type Math struct{}

func (m *Math) Mean(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}

func (m *Math) Median(nums []int) float64 {
	sort.Ints(nums)
	mid := len(nums) / 2
	if len(nums) == 0 {
		return 0
	}
	if len(nums)%2 != 0 {
		return float64(nums[mid])
	}
	return float64(nums[mid-1]) + float64(nums[mid])/2
}

// This can calulate the frequency of an element
func (m *Math) Mode(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	frequency := make(map[int]int)
	maxFreq := 0
	var modes []int
	for _, v := range nums {
		frequency[v]++
		if frequency[v] > maxFreq {
			maxFreq = frequency[v]
		}
	}

	for i, val := range frequency {
		if val == maxFreq {
			modes = append(modes, i)
		}
	}
	return modes
}

// calculate the range value
func (m *Math) Range(nums []int) int {
	max := nums[0]
	min := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}
	return max - min
}

func (m *Math) Max(nums []int) int {
	max := nums[0]
	for i := range nums[:] {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func (m *Math) Min(nums []int) (int, error) {
	// [2,3,4,1]
	min := nums[0]
	for i := range nums[1:] {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min, nil
}

func (m *Math) SquareRoot(n int) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("square root of a negative number is undefined")
	}

	if n == 0 || n == 1 {
		return float64(n), nil
	}

	// Convert `int` to `float64` for calculation
	x := float64(n)
	tolerance := 1e-7 // Precision

	for {
		nextX := 0.5 * (x + float64(n)/x) // Newton-Raphson formula

		// Check convergence manually (abs without math package)
		diff := nextX - x
		if diff < 0 {
			diff = -diff
		}

		if diff < tolerance {
			return nextX, nil
		}

		x = nextX
	}
}

func (m *Math) Power(n int, exp int) int {
	// 2^3= 2*2*2
	val := 1
	if exp == 0 {
		return 1
	}
	for i := 0; i < exp; i++ {
		val *= n
	}
	return val
}
