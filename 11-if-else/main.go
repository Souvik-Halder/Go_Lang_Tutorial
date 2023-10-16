package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on if else")

	logincount:= 10


	var result string
	if logincount < 10 { //here in that line only the braces should start
		result = "Regular User"
	}else if logincount > 10 { //here in that line only the else if should start
		result = "Watch out"
	}else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)	

	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else {
		fmt.Println("Number is odd")
	}

	if num:=3; num<10 { //here we can also declare a variable inside the if statement
		fmt.Println("Number is less than 10")
	}else {
		fmt.Println("Number is not less than 10")
	}
	 
	

}