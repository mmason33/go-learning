package main

import "fmt"

// individually typed params
// unnamed typed return value
func add(x int, y int) int {
	return x + y
}

// implicitly returns value for "z"
func subtract(x, y int) (z int) {
	z = x - y
	return
}

// multiple named return values
// named or multiple return types/values require parens
func multiplyDivide(x, y int) (product, quotient int) {
	product = x * y
	quotient = x / y
	return
}

// clearer way - NO NAMING OF RETURN VALUES
func multiplyDivideTwo(x, y int) (int, int) {
	product, quotient := x*y, x/y
	return product, quotient
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(subtract(2, 1))
	fmt.Println(multiplyDivide(4, 2))
	fmt.Println(multiplyDivideTwo(4, 2))
}
