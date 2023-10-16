package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to a class on switch case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 //rand.Intn(6) will give a random number between 0 to 5 and we are adding 1 to it to get a random number between 1 to 6
	fmt.Println("Value of dice is: ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough //fallthrough is used to execute the next case also
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 3 spots")
	case 5:
		fmt.Println("You can move 3 spots")
	case 6:
		fmt.Println("You can move 3 spots and roll the dice again")
	}
}
