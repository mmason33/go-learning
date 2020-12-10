package main

import "fmt"

func fullFor() {
	sum := 0 // shorthand
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

func forWithoutPrePost() {
	// For loop without pre/post statements
	var sum int = 1 // longhand
	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)
}

func forWhile() {
	// For loop as a while loop
	sum := 1 // shorthand
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

func main() {
	fullFor()
	forWithoutPrePost()
	forWhile()
}
