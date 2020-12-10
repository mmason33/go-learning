package main

import "fmt"

func main() {
	s := []string{}
	// append value to end of slice
	s = append(s, "hello")
	// apprend multiple values to end of slice
	s = append(s, "world", "!")
	fmt.Println(s)

	// append slice to another slice
	t := []string{"hey"}
	r := []string{"there"}
	// you can also append a slice to another using an ellipsis
	// using the ellipsis (...) after our slice, we indicate that we want to append each element of our slice
	// types need to match
	t = append(t, r...)
	fmt.Println(t)

	// len => lenght of slices
	cities := []string{
		"Santa Monica",
		"San Diego",
		"San Francisco",
	}
	fmt.Println(len(cities))
	// 3
	countries := make([]string, 42)
	fmt.Println(len(countries))

	var z []string // this only works as a "nil" slice because the slice is declared without a value
	var y []string = []string{}
	fmt.Println(y, len(y), cap(y))
	fmt.Println(z, len(z), cap(z))
	// [] 0 0
	if z == nil {
		fmt.Println("z nil!")
	}

	if y != nil {
		fmt.Println("y not nil!")
	}
}
