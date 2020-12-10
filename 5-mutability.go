package main

import "fmt"

// struct custom "type"
type Artist struct {
	Name, Genre string
	Songs       int
}

// increments songs by value
// arguments are passed by value, a function receiving a value argument and mutating it, won’t mutate the original value.
func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

// "*" struct essentially processes the address and retrieves the value of an address pointer
func newReleasePoint(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	// declare me with a type Artist
	me := Artist{Name: "Non-Pointer", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)

	// declare mePointer as a pointer to a struct of type Artist
	mePointer := &Artist{Name: "Pointer", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", mePointer.Name, newReleasePoint(mePointer))
	fmt.Printf("%s has a total of %d songs\n", mePointer.Name, mePointer.Songs)
	fmt.Println(100 << 1)
}
