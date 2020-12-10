package main

import "fmt"

var a [10]int // array of 10 integers - array's length is part of it's type

func main() {
	// zerovalued array declaration
	fmt.Println(a)

	// Explicit value by index
	var b [2]string
	b[0] = "Hello"
	b[1] = "World"
	fmt.Println(b[0], b[1])
	fmt.Println(b)

	// Value with length declaration
	c := [2]string{"hello", "world!"}
	fmt.Printf("%q", c)

	// Implicit/inferred length declaration
	d := [...]string{"hello", "world!"}
	fmt.Printf("%q", d)
}
