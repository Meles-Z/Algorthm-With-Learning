package main

import "fmt"

func main() {
	var numbers int
	var store []int

	for range 10 {
		fmt.Print("Enter a number: ")
		fmt.Scan(&numbers)
		store = append(store, numbers)
	}

	var t string
	fmt.Print("Do you want positive or negative numbers? (p/n): ")
	fmt.Scan(&t)

	var filtered []int
	for _, num := range store {
		if t == "p" && num > 0 {
			filtered = append(filtered, num)
		} else if t == "n" && num < 0 {
			filtered = append(filtered, num)
		}
	}

	if len(filtered) > 0 {
		fmt.Println("Filtered numbers:", filtered)
	} else {
		fmt.Println("No numbers matched your selection.")
	}
}
