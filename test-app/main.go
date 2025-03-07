package main

import "fmt"

func main() {
	test := [][]int{{1, 2, 3}, {4, 5, 6}}
	for _, v := range test {
		for _, t := range v {
			fmt.Println(t)
		}
	}
}
