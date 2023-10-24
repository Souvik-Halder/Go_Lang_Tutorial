package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("This is the first line")// This defer line will be executed at the end of the main function
defer fmt.Println("One")
defer fmt.Println("Two")
	fmt.Println("Welcome to defer in golang")
	myDefer()
}

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}

}