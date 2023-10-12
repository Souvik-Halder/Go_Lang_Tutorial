package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on structs")
	//no inheritance in golang : no super or parent keyword
	hitesh := User{"Souvik", "souvik.gmail.com", 21, true}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh) //It will print the details of the hitesh struct//It will print the all fields and the values of the struct
	fmt.Println("Name is: ", hitesh.Name, "Email is: ", hitesh.Email, "Age is: ", hitesh.Age, "Status is: ", hitesh.Status)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

//User should be in capital letter as it is the struct this is a common practice of nomenculture in golang
