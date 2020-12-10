package main

import "fmt"

/*
TASK
Copy your Sqrt function from the earlier exercises (Section 5.4) and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn’t support complex numbers.
*/

// alias
type ErrNegativeSqrt float64

// alias implements built-in error interface
// implicit implementation by defined Error method required by interface type error
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-f)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(-2.0))
	fmt.Println(Sqrt(16))
	// Creating a value of type ErrNegativeSqrt will automatically run the Error method because it implements type error interface
	fmt.Println(ErrNegativeSqrt(2.))
}
