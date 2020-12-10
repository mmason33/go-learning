package main

import "fmt"

/*
*
* Basically the return object of a constructor function/class called Human
*
*
class Human {
constructor(Alive) {
	this.Alive = Alive
}
}
*
*/
type Human struct {
	Alive bool
}

/*
class Person extends Human {
	constructor(args) {
		super(args.Alive)
		this.Name = args.Name
		this.Age = args.Age
	}
}
*/
type Person struct {
	Human
	Name string
	Age  int
}

/*
interface Score {
	GetScore()
}
*/
type Score interface {
	GetScore() int
}

/*
class Player extends Person {
	constructor(args) {
		super(args)
		this.Score = args.Score
	}
}
*/
type Player struct {
	Person
	Score int
}

/*
This is essentially implicitly saying - class Player implements interface Score
Without this method error would be thrown
Error: Player does not implement Score (missing GetScore method)
*/
func (p Player) GetScore() int {
	return p.Score
}

/*
ShowScore takes in a struct that implements interface Score
*/
func ShowScore(s Score) int {
	return s.GetScore()
}

func main() {
	// Shorthand struct literal
	player := Player{
		Person{
			Human{
				true,
			},
			"Mike",
			30,
		},
		9000,
	}

	// Longhand struct literal
	player2 := Player{
		Person: Person{
			Human: Human{
				Alive: true,
			},
			Name: "Henry",
			Age:  42,
		},
		Score: 2000,
	}

	fmt.Println(player)
	fmt.Println(player2)
	fmt.Println("---------------")
	fmt.Println(player.Person.Name, "=>", ShowScore(player))
	fmt.Println(player2.Person.Name, "=>", ShowScore(player2))
	// {{{true} Mike 30} 9000}
	// {{{true} Henry 42} 2000}
	// ---------------
	// Mike => 9000
	// Henry => 2000
}
