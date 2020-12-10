package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

// (u *User) - is the receiver of this method
// Method will be called on a pointer where the location holds a struct of type User
func (u *User) Greeting() string {
	// Properties FirstName and LastName are available because this method must be call on a pointer whose location contains a struct of type User
	// (u *User) - dereference pointer value "*" - the actual value of a pointer location
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := &User{"Guy", "Finkle"}
	fmt.Println(u.Greeting())
}
