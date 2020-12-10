package main

import (
	"fmt"
	"reflect"
)

func basicIf(s string) bool {
	var res bool = false
	if reflect.TypeOf(s).String() == "string" {
		res = true
	}

	return res
}

func basicIfElse(m map[string]string) bool {
	var res bool
	var length int = len(m)
	if length > 2 {
		res = true
	} else {
		res = false
	}

	return res
}

func basicIfElseIfElse(num int) string {
	var str string
	if num > 200 {
		str = "num greater than 200"
	} else if num > 100 {
		str = "num greater than 100 less than 200"
	} else {
		str = "num less than 100"
	}

	return str
}

func main() {
	fmt.Println(basicIf("test"))
	fmt.Println(basicIfElse(map[string]string{"1": "1", "2": "3", "3": "4"}))
	fmt.Println(basicIfElse(map[string]string{"1": "1", "2": "3"}))
	fmt.Println(basicIfElseIfElse(201))
	fmt.Println(basicIfElseIfElse(101))
	fmt.Println(basicIfElseIfElse(99))
}
