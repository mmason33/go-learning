package main

import "fmt"

// A slice points to an array of values and also includes a length. Slices can be resized since they are just a wrapper on top of another data structure.
func main() {
	mySlice := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(mySlice)
	// [2 3 5 7 11 13]

	fmt.Println(mySlice[1:4])
	// [3 5 7]

	// missing low index implies 0
	fmt.Println(mySlice[:3])
	// [2 3 5]

	// missing high index implies len(s)
	fmt.Println(mySlice[4:])
	// [11 13]
}
