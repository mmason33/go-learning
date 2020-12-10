package main

import "fmt"

func main() {
	m := map[string]string{
		"name":     "Mike",
		"job":      "Engineer",
		"location": "California",
	}

	// key vs index
	for key, value := range m {
		fmt.Println(key, "=>", value)
	}
}
