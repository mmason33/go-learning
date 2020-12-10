package main

import "fmt"

/*
	Array - fixed size defined at declaration - numerically indexed values - can be before/after initialization(given a value)
	Slice - mutable length - items can be added or removed - numerically indexed values - make method to declare a slice - append method to add value to a slice(end)
	Map - mutable length - keys are of a specified type(string int) - make method to declare map
*/

// Declare without initialization(value)
var (
	arr [5]int         // zerovalue => [0,0,0,0,0]
	slc []int          // zerovalue => [] => nil
	mp  map[string]int // zerovalue => map[] => nil
)

func main() {
	fmt.Println(arr)
	fmt.Println(slc, slc == nil) // true
	fmt.Println(mp, mp == nil)   // true
	fmt.Println("--------------------")

	// shorthand from with function declaration

	// array
	a := [3]int{1, 2, 3}

	// slice
	s := []int{1, 2, 3}

	// map
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3, // trailing comma needed
	}

	ms := make([]string, 5)

	mm := make(map[string]string)

	fmt.Println(a, s, m, ms, mm)
}
