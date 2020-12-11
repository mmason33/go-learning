package main

import (
	"fmt"
	"time"
)

/*
GOROUTINES - A goroutine is a lightweight thread managed by the Go runtime. - CONCURRENCY
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	say("1")
	say("2")
}
