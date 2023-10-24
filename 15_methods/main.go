package main

import "fmt"

type User struct{
	Name string
	Email string
	Status bool
	Age int
}


func main(){
	souvik:= User{"Souvik", "souvik@gmail.com", true, 25}	
	souvik.GetStatus()
	souvik.NewMail()
	souvik.GetStatus()
	fmt.Println("Welcome to methods in golang")
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
	fmt.Println("User name is: ", u.Name)
	fmt.Println("User email is: ", u.Email)

}

func (u User) NewMail(){
	u.Email="test@godev.in"
	fmt.Println("User email is: ", u.Email)//here copy of the object is passed not the actual object
}