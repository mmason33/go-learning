package main

import "fmt"

func main() {
	/*
		DECLARED VARIABLES WITHOUT INITIALIZATION
	*/
	var num int
	fmt.Println(num) // 0

	var str string
	fmt.Println(str) // ""

	var arr [0]int
	fmt.Println(arr) // []

	// zerovalue for pointers, functions, interfaces, slices, channels, and maps is nil
	// looks like an empty array but it's actually nil
	var slice []string
	fmt.Println(slice)
	fmt.Println("slice equal to nil", slice == nil)

	// Does have initialization
	// Initialized to empty slice of type string
	var slice2 []string = []string{}
	fmt.Println(slice2)
	fmt.Println("slice equal to nil", slice2 == nil)
}
