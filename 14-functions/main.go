package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(2,3)
	fmt.Println(result,"This is the addition of two numbers")

	proResult,myMessage :=proAdder(2,3,4,221)
	fmt.Println("The proresult is ",proResult,myMessage)


}

func adder(val1 int ,val2 int) int{
	result := val1+val2
	return result
}

func proAdder(values ...int) (int,string){
	total :=0
	for _,val := range values{
		total+=val
	}
	return total,"Hi this function returns multiple things"
}

func greeter(){
	fmt.Println("Good morning")
}
