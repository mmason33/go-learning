package main

import "fmt"

type test_struct struct {
	Message string
}

func (t test_struct) Say() {
	fmt.Println(t.Message)
}

func (t test_struct) Update(m string) {
	t.Message = m
}

func (t *test_struct) SayP() {
	fmt.Println(t.Message)
}

func (t *test_struct) UpdateP(m string) {
	t.Message = m
}

func main() {
	// Each time a copy of the caller is created and passed to the function
	// np (receiver) is copied before each method call
	np := test_struct{"test"}
	np.Say() // prints test
	np.Update("test updated")
	np.Say() // prints test

	// When you pass a pointer as an argument, what happens under the hood is that a copy of that pointer is created and passed to the underlying function
	// So you can modify the value of the caller
	p := &test_struct{"testP"}
	p.SayP() // testP
	p.UpdateP("testP updated")
	p.SayP() // testP updated
}
