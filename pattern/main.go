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

func main() {
	// decreasePattern(5)
	trianglePattern(4)
}
