// Go supports the new expression to allocate a zeroed value of the requested type and to return a pointer to it.
// x := new(int) returns a pointer to the zerovalued type int
// to get the value of the pointer x you need to dereference the pointer
// fmt.Println(*x)
package main

import "fmt"

type Location struct {
	Lat  float64
	Long float64
}

func main() {
	// Initializing a struct using "new" will set properties to their zerovalued value of the specified type => float64 zerovalue is 0
	// New is constructor same idea as JavaScript
	lat := new(Location)
	// Declaring a pointer to a defined/declared "struct literal"
	// Struct literal is equivilant to an object literal in JavaScript
	// In this case below is a pointer to a struct literal
	/*
		long := Location{} // struct of type Location in Go
		let Location = {} // type object in JS
	*/
	long := &Location{}
	fmt.Println(*lat == *long) // compare the dereference values of pointers to zerovalued structs => result is true they are equal
}
