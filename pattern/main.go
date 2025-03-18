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
	num := doubleNumbers(nums)
	fmt.Println("Sum:", num)
	return num/10 == 0
}

func doubleNumbers(n int) int {
	double := n * 2
	sum := 0
	if double > 10 {
		sum = 1 + n%10
	} else {
		sum = double
	}
	fmt.Println("doubled number:", double)
	return sum
}

func main() {
	// decreasePattern(5)
	// trianglePattern(4)
	fmt.Println(doubleNumbers(349))
	fmt.Println(isLuhan(349))
}
