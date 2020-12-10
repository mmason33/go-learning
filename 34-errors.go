package main

import (
	"fmt"
	"reflect"
	"time"
)

// Error struct
type MyError struct {
	When time.Time
	What string
}

// Built-in error interface
// type error interface {
//     Error() string
// }

// Implements built-in error interface by declaring Error method on the receiver - pointer to a struct of type MyError
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() *MyError { // different return types
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func run2() error { // different return types
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// same thing for some reason
	// both are type *MyError
	if err := run(); err != nil {
		fmt.Println(err)
		fmt.Println(reflect.TypeOf(err))
	}

	if err2 := run2(); err2 != nil {
		fmt.Println(err2)
		fmt.Println(reflect.TypeOf(err2))
	}
}
