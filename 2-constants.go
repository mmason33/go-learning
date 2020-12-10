package main

// std-lib I/O
// singular import
import "fmt"

// inferred multi const declaration
const (
	Pi    = 3.14
	Truth = false
	Big   = 1 << 62
	Small = Big >> 61
)

func main() {
	const Greeting = "ハローワールド"
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)
}
