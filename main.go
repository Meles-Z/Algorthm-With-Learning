package main

import (
	"fmt"
)

type Person struct {
	name  string
	email string
}
type myint int
type mystring string

// when i use anonays function i can use type like below
type myfunc func(n int)

func increment(n int) {
	fmt.Printf("The incremented Value is:%d", n+1)
}

func square(n int) {
	fmt.Printf("The squeared value is:%d", n*n)
}

func main() {
	var f myint = 6
	var t mystring = "Hi there!"
	fmt.Println(t)
	fmt.Println(f)
	p := Person{
		name:  "s",
		email: "m",
	}
	fmt.Printf("%+v\n", p)

	
	var inc myfunc=increment
	fmt.Println(inc)
	inc=square
	fmt.Println(inc)
}
