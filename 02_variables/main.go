package main

import "fmt"

const LoginToken string = "kjfhflksdflksd" //constant variable

func main() {

	var userName string = "Souvik Halder"
	//%T will print the type of the variable

	fmt.Printf("Variable is of type: %T \n", userName)
	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	var smallVal uint8 = 255
	fmt.Printf("Variable is of type: %T \n", smallVal)
	fmt.Println(smallVal)

	//float32 contains only the 5 decimal places and float 64 contains 15 decimal places

	var smallFloat float32 = 255.4555555
	fmt.Println(smallFloat)

	//default values and some aliases
	var website = "hello"
	fmt.Println(website)
	//here lexer comes in and it will automatically detect the type of the variable and it also detect the semicolon

	numberOfUser := 300000 //This is valid
	//inside a method we can use := but outside a method we can't use := this operator

	fmt.Println(numberOfUser)

	//In go lang u need to declare the global variable by capital Letter and local variable by small letter this is the official syntax
	//as here the LoginToken is the global variable
	fmt.Println(LoginToken)

}
