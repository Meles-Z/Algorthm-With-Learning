// package main

// import (
// 	"fmt"
// )

// type Person struct {
// 	name  string
// 	email string
// }
// type myint int
// type mystring string

// // when i use anonays function i can use type like below
// type myfunc func(n int)

// func increment(n int) {
// 	fmt.Printf("The incremented Value is:%d", n+1)
// }

// func square(n int) {

// 	fmt.Printf("The squeared value is:%d", n*n)
// }

// func main() {
// 	var f myint = 6
// 	var t mystring = "Hi there!"
// 	fmt.Println(t)
// 	fmt.Println(f)
// 	p := Person{
// 		name:  "s",
// 		email: "m",
// 	}
// 	fmt.Printf("%+v\n", p)

// 	var inc myfunc=increment
// 	fmt.Println(inc)
// 	inc=square
// 	fmt.Println(inc)
// }

// package main

// import "fmt"

// func lengthOfLongestSubstring(s string) int {
// 	// Map to store the last seen index of characters
// 	lastSeen := make(map[rune]int)
// 	start, maxLen := 0, 0

// 	for i, char := range s {
// 		// If the character is found in the map and is within the current window
// 		if lastIdx, found := lastSeen[char]; found && lastIdx >= start {
// 			start = lastIdx + 1
// 		}
// 		lastSeen[char] = i // Update last seen index
// 		currentLen := i - start + 1
// 		if currentLen > maxLen {
// 			maxLen = currentLen
// 		}
// 	}

// 	return maxLen
// }

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
// 	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
// }

package main

import "fmt"

func main() {
	// your goal
	// restate problem clearly
	// devide problem
	// start with what you know
	// optimize what you did

	// var num int
	// var store []int

	// for range 10 {
	// 	fmt.Print("Enter the number:")
	// 	fmt.Scan(&num)
	// 	store = append(store, num)
	// }

	// var t string
	// fmt.Print("Do you want possitive or negative number (p/n) ?")
	// fmt.Scan(&t)
	// if t == "p" && num > 0 {
	// 	fmt.Println(store)
	// } else if t == "n" && num < 10 {
	// 	fmt.Printf("%d\n", store)
	// } else {

	// 	fmt.Println("Invalid input")
	// }

	// for i:=1; i<5; i++{
	// 	for j:=0; j<i; j++{
	// 		fmt.Print("* ")
	// 	}
	// 	fmt.Println()
	// }

	// for i:=4; i>0; i--{
	// 	for j:=1; j<i; j++{
	// 		fmt.Print("* ")
	// 	}
	// 	fmt.Println()
	// }

	// start with what you know
	// rows := 8
	// // maxWidth := rows
	// for i := rows; i > 1; i-- {
	// 	if i%2 == 0 {
	// 		space := (rows - i)/2
	// 		for s := 0; s < space; s++ {
	// 			fmt.Print(" ")
	// 		}
	// 		for j := 0; j < i; j++ {
	// 			fmt.Print("#")
	// 		}

	// 		fmt.Println()
	// 	}
	// }

	// target := 4
	// arr := []int{1, 2, 3, 4, 5}
	// for i, row := range arr {
	// 	if row == target {
	// 		fmt.Println("Found at index:", i)
	// 		break
	// 	}
	// }

	rows := []int{100, 50, 4, 80, 60}
	// greaterValue := rows[0]
	// for i := 1; i < len(rows); i++ {
	// 	if rows[i]>greaterValue{
	// 		greaterValue = rows[i]
	// 	}
	// }
	// fmt.Println(greaterValue)

	
}
