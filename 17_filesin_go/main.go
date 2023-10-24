package main

import (
	"io/ioutil"
	"io"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Welcome to files in golang")
	content:= "This is a test file"
	file,err:=os.Create("../test.txt")
	checkNillError(err)
	length,err := io.WriteString(file,content)//this will return the length of the string written in the file also this will write in the file
	checkNillError(err)//This is the previous error syntax but here the code is in the function named as checkNillError
	fmt.Println("Length is: ",length)
	readFile("../test.txt")
}

func readFile(fileName string){
	databyte,err := ioutil.ReadFile(fileName)
	checkNillError(err)
	fmt.Println("Data as bytes: ",databyte)//This will print the data as bytes
	fmt.Println("Data as string: ",string(databyte))//This will print the data as string

}

func checkNillError(err error){
	if err!=nil{
		
		panic(err)//panic is used to stop the execution of the program
	}
}