package main

// https://codeburst.io/diving-deep-into-the-golang-channels-549fd4ed21a8
import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// func goRoutineA(a <-chan int) {
// 	val := <-a
// 	fmt.Println("goRoutineA received the data", val)
// }

// func goRoutineB(a <-chan int) {
// 	val := <-a
// 	fmt.Println("goRoutineB received the data", val)
// }
// func main() {
// 	ch := make(chan int)
// 	go goRoutineA(ch)
// 	go goRoutineB(ch)
// 	x, y := <-ch, <-ch
// 	fmt.Println(x, y)
// 	// time.Sleep(time.Second * 1)
// }
