package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var min int
	var max int

	fmt.Println("Choose a minimum value")
	fmt.Scan(&min)
	fmt.Println("Choose a maximum value")
	fmt.Scan((&max))

	rand.Seed(time.Now().UnixNano()) //makes r is different each time

	r := rand.Intn(max-min+1) + min
	loop := 0

	for loop < 1 {
		fmt.Println("Enter a guess")

		var guess int
		fmt.Scan(&guess)
		if guess == r {
			fmt.Println("You win")
			loop = 1
		} else if guess > r {
			fmt.Println("Too high, try again")
		} else if guess < r {
			fmt.Println("Too low, better luck next time")
		}
	}
}
