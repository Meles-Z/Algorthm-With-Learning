package main

import "fmt"

func decreasePattern(n int) {
	for i := n; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func trianglePattern(n int) {
	//n=5
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func isLuhan(nums int) bool {
	num := sumOfDigit(nums)
	if num/10 != 0 {
		return false
	}
	return true
}
func sumOfDigit(n int) int {
	store := []int{}
	for n != 0 {
		singleNumber := n % 10
		store = append(store, singleNumber)
		n /= 10
	}
	sum := doubleNumbers(n)
	for _, v := range store {
		sum += v
	}
	return sum
}
func doubleNumbers(n int) int {
	double := n * 2
	sum := 0
	if double > 10 {
		sum = 1 + n%10
	} else {
		sum = double
	}
	return sum
}

func main() {
	// decreasePattern(5)
	// trianglePattern(4)
	fmt.Println(doubleNumbers(22))
	fmt.Println(isLuhan(345))
}
