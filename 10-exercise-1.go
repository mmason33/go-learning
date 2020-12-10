package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

// passed by reference
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

// passed by value
func (u User) Goodbye() string {
	return fmt.Sprintf("Goodbye %s from %s",
		u.Name, u.Location)
}

type Player struct {
	*User  // pointer to a User struct
	GameId int
}

func main() {
	// insert code
	p := Player{ // composite/struct literal
		User: &User{
			Id:       11,
			Name:     "Exercise",
			Location: "localhost",
		},
		GameId: 22,
	}

	// Call Greetings on the receiver of type => *User = dereference of a pointer aka the value(User struct) of a pointer slot
	fmt.Println(p.Greetings()) // Hi Exercise from localhost
	fmt.Println(p.Goodbye())   // Goodbye Exercise from localhost
	fmt.Println(p)             // {0xc000062180 22}
	fmt.Println(*p.User)       // {11 Exercise localhost}
}
