package main

import (
	"fmt"
	"strings"
)

// Custom type
type StringProxy string

type SliceProxy []string

func (s StringProxy) UpperCase() string {
	// strings.ToUpper(s) does not work because ToUpper expects a string and "s" is a type of StringProxy
	return strings.ToUpper(string(s))
}

func main() {
	// Declare custom type - similar to declaring struct literal
	// Primative/basic types(strings, ints, bools, floats) are declared with () vs {} from structs/maps/slices/arrays
	fmt.Println(StringProxy("test"))
	fmt.Println(StringProxy("test").UpperCase())
	fmt.Println(SliceProxy{"hey", "there"})
}
