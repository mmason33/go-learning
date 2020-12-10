package main

import (
	"strings" // had hint - You might find strings.Fields helpful. - https://golang.org/pkg/strings/#Fields

	"golang.org/x/tour/wc" // The wc.Test function runs a test suite against the provided function and prints success or failure
)

// TASK: Implement WordCount.
func WordCount(s string) map[string]int {
	// Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
	t := strings.Fields(s)    //
	m := make(map[string]int) // declare map without value
	for _, value := range t {
		m[value]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
