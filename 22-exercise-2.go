package main

import "fmt"

// TASK: Given a list of names, you need to organize each name within a slice based on its length.
var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

// solution
func main() {
	var maxLen int // length holder

	for _, name := range names {
		if l := len(name); l > maxLen { // assigns variable "l" to the length of name for a given iteration and compares to the persisting value of maxLen
			maxLen = l
		}
	}

	output := make([][]string, maxLen) // declare a slice containing slices containing strings - length is maxLen in this case 9
	test := make([]string, 5)
	// for _, name := range names {
	// 	output[len(name)-1] = append(output[len(name)-1], name)
	// }

	fmt.Printf("%v", output)
}

// func main() {
// one := make([]string, 0)
// two := make([]string, 0)
// three := make([]string, 0)
// four := make([]string, 0)
// five := make([]string, 0)
// six := make([]string, 0)
// seven := make([]string, 0)
// eight := make([]string, 0)
// nine := make([]string, 0)
// for _, value := range names {
// 	length := len(value)
// 	switch {
// 	case length == 1:
// 		one = append(one, value)
// 	case length == 2:
// 		two = append(two, value)
// 	case length == 3:
// 		three = append(three, value)
// 	case length == 4:
// 		four = append(four, value)
// 	case length == 5:
// 		five = append(five, value)
// 	case length == 6:
// 		six = append(six, value)
// 	case length == 7:
// 		seven = append(seven, value)
// 	case length == 8:
// 		eight = append(eight, value)
// 	case length == 9:
// 		nine = append(nine, value)
// 	}
// }

// results := [][]string{
// 	one,
// 	two,
// 	three,
// 	four,
// 	five,
// 	six,
// 	seven,
// 	eight,
// 	nine,
// }

// fmt.Println(results)
// }
