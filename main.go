package main

import "fmt"

// create a function that take string as an input, and return it the revese string
func reverseString(s string) string {
	// create a variable that can store this value
	reverse := ""
	// start the loop from end
	for i := len(s) - 1; i >= 0; i-- {
		// store the loop value of each index into the value
		reverse += string(s[i])
	}
	// return the stored valued in the funcion
	return reverse
}

func main() {
	fmt.Println(reverseString("yung"))
}
