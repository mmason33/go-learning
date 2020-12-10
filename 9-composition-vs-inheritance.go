package main

import (
	"fmt"
	"log"
	"os"
)

// Struct composition
// No inheritance in Go
type User struct {
	Name, Location string
	Id             int
}

// User struct method available on Player struct
/*
	(u *User) => This is called the receiver
	The object in which the method will be called on
*/
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

type Player struct {
	User   // no typing on User field
	gameId int
}

// Second example
type Job struct {
	Command string
	*log.Logger
}

func main() {
	// Dot notation composition
	p := Player{}
	p.Id = 33
	p.Name = "Guy"
	p.Location = "California"
	p.gameId = 203
	fmt.Println(p) // does not print keys {{Guy California 33} 203}
	fmt.Println(p.Greetings())

	// Struct literal composition
	p2 := Player{
		User: User{
			Name:     "Guy2",
			Location: "California2",
			Id:       34, // trailing comma
		}, // trailing comma
		gameId: 204, // trailing comma
	}
	fmt.Println(p2) // {{Guy2 California2 34} 204}

	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Print("starting now...")
}
