package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")
	var ptr *int //it will give the nill value
	fmt.Println("The value of pointer is: ", ptr)
	myNumber := 23

	var ptr1 = &myNumber
	fmt.Println("The memory address of pointer is: ", ptr1)
	fmt.Println("The value of pointer is: ", *ptr1) // * is used to get the value of the pointer
	*ptr1 = *ptr1 + 2
	fmt.Println("The newvalue of pointer is: ", *ptr1)
}
