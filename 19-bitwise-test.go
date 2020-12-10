package main

import "fmt"

/*
BITWISE SHIFT OPERATORS
Literally shifts the bits for a/b to the left or right n times = a/b << n
Makes alot more sense when using int8 vs uint or int 32 or 64 bit integers
*/

func main() {
	var a int8 = 3
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", a<<1)
	fmt.Printf("%08b\n", a<<2)
	fmt.Printf("%08b\n", a<<3)

	var b uint8 = 120
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", b>>1)
	fmt.Printf("%08b\n", b>>2)
}
