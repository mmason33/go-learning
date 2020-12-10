// A struct is a collection of fields/properties. You can define new types as structs or interfaces
package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	// Latitude of the event
	Lat float64
	// Longitude of the event
	Lon float64
	// Date of the event
	Date time.Time
}

type Point struct {
	x, y int
}

type Car struct {
	Make, Model, Trim string
}

func main() {
	// All of whats below is:
	// Declaring a struct literal or
	// Declaring a pointer to struct literal
	/*
		point := Location{} // struct of type Location in Go
		let point = {} // type object in JS
	*/
	fmt.Println(
		Bootcamp{
			Lat:  34.012836,
			Lon:  -118.495338,
			Date: time.Now(),
		})

	p := Point{1, 2}  // type point implicit x y values
	q := &Point{1, 2} // Pointer to a struct with implicit x y values
	r := Point{x: 1}  // explicit x value and implicit y: 0
	s := Point{}      // implicit x: 0 and y: 0
	fmt.Println(p, q, r, s)

	car := Car{
		Make:  "Dodge",
		Model: "Charger",
		Trim:  "Hellcat", // always need trailing comma
	}
	fmt.Println(car.Make, car.Model, car.Trim) // struct dot notation
}
