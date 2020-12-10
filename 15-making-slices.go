package main

import "fmt"

// Besides creating slices by passing the values right away (slice literal), you can also use make.
func main() {
	cities := make([]string, 3)
	cities[0] = "Santa Monica"
	cities[1] = "Venice"
	cities[2] = "Los Angeles"
	fmt.Printf("%q", cities)
}
