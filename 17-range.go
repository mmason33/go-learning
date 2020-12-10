package main

import "fmt"

var pow []int = []int{1, 2, 4, 8, 16, 32, 64, 128}
var alphabet []string = []string{"a", "b", "c", "d", "e"}

func main() {
	// index, value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// skip index or value by assigning it to "_"
	for _, val := range alphabet {
		fmt.Println(val)
	}

	// break range
	p := make([]int, 10) // slice length 10 containing zeros
	for i := range p {
		p[i] = 1 << uint(i)
		if p[i] >= 16 {
			break
		}
	}
	fmt.Println(p)

	// continue range
	pw := make([]int, 10) // slice length 10 containing zeros
	for i := range pw {
		if i%2 == 0 {
			continue
		}
		pw[i] = 1 << uint(i)
	}
	fmt.Println(pw)
}
