package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 map[string]Vertex

func main() {
	// using make to give m a value
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)

	// defined a map literal
	m2 = map[string]Vertex{
		// "Google": Vertex{37.42202, -122.08408}, same thing
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m2)

	// Map mutations and delete
	m3 := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	m3["4"] = 4
	elm := m3["4"]
	delete(m3, "3") // delete(map, key)
	fmt.Println(m3, elm)
}
