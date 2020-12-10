package main

/*
TASK
You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth The coins will be distributed based on the vowels contained in each name where:

a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins

and a user can’t get more than 10 coins. Print a map with each user’s name and the amount of coins distributed. After distributing all the coins, you should have 2 coins left.
*/

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

// my working solution
// func main() {
// 	for _, name := range users {
// 		nameSlice := strings.Split(name, "")
// 		distribution[name] = 0

// 		for _, letter := range nameSlice {
// 			switch letter {
// 			case "a", "A", "e", "E":
// 				if distribution[name]+1 > 10 {
// 					amount := 10 - distribution[name]
// 					distribution[name] += amount
// 					coins -= amount
// 					break
// 				}
// 				distribution[name]++
// 				coins--
// 			case "i":
// 				if distribution[name]+2 > 10 {
// 					amount := 10 - distribution[name]
// 					distribution[name] += amount
// 					coins -= amount
// 					break
// 				}
// 				distribution[name] += 2
// 				coins -= 2
// 			case "o":
// 				if distribution[name]+3 > 10 {
// 					amount := 10 - distribution[name]
// 					distribution[name] += amount
// 					coins -= amount
// 					break
// 				}
// 				distribution[name] += 3
// 				coins -= 3
// 			case "u":
// 				if distribution[name]+4 > 10 {
// 					amount := 10 - distribution[name]
// 					distribution[name] += amount
// 					coins -= amount
// 					break
// 				}
// 				distribution[name] += 4
// 				coins -= 4
// 			}
// 		}
// 	}

// 	fmt.Println(distribution)
// 	fmt.Println("Coins left:", coins)
// }

// The given solution
// way better cleaner solution
func main() {
	coinsForUser := func(name string) int {
		var total int
		for i := 0; i < len(name); i++ {
			switch string(name[i]) {
			case "a", "A":
				total++
			case "e", "E":
				total++
			case "i", "I":
				total = total + 2
			case "o", "O":
				total = total + 3
			case "u", "U":
				total = total + 4
			}
		}
		return total
	}

	for _, name := range users {
		v := coinsForUser(name)
		fmt.Println(v)
		if v > 10 {
			v = 10
		}
		distribution[name] = v
		coins = coins - v
	}
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
