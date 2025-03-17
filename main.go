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

func maxValue(nums []int) int {
	max := nums[0]
	for i := range nums[1:] {
		if nums[i+1] > max {
			max = nums[i+1]
		}
	}
	return max
}

var lorem int =3
func init(){
	lorem=5
}
func main() {
	// fmt.Println(maxValue([]int{2, 3, 10, 4, 4, 5, 6}))
	fmt.Println(lorem)
}
