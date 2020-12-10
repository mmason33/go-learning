package main

import (
	"fmt"
)

type Sayer interface {
	Speak()
}

type Sentence struct {
	Message string
}

func SaySentence(s Sayer) {
	s.Speak()
}

func (s Sentence) Speak() {
	fmt.Println(s.Message)
}

func (s Sentence) Listen() {
	fmt.Println("Listening")
}

func main() {
	s := Sentence{
		Message: "Hello world",
	}

	SaySentence(s)
	s.Listen()
}
