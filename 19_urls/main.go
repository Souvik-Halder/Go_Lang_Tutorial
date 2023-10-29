package main

import (
	"fmt"
	"net/url"
)

const myurl string ="https://lco.dev:3000/learn?course=reactjs&paymentid=123"

func main() {
fmt.Println("Welcome to urls in golang")
fmt.Println(myurl)

//parsing the url
   result,_ := url.Parse(myurl)
   
   fmt.Println(result.Scheme)//This means protocol of the url
   fmt.Println(result.Host)//This means host of the url
   fmt.Println(result.Path)//This means path of the url
   fmt.Println(result.Port())//This means port of the url
   fmt.Println(result.RawQuery)//This means query of the url
   fmt.Println(result.Query())//This means query of the url in the form of map
   qParams := result.Query()	
   fmt.Printf("Type of qParams is: %T\n",qParams)
   fmt.Println(qParams["course"])//This will print the value of the key course
   fmt.Println(qParams["paymentid"])//This will print the value of the key paymentid

   //By using for loop
   for _,val:= range qParams{
	   fmt.Println("param is",val)
   }

   partsOfUrl := &url.URL{
	   Scheme: "https",
	   Host: "lco.dev",
	   Path: "/tutcss",
	   RawPath: "user=souvik",
   }

   anotherUrl:=partsOfUrl.String()//This another way to handle the string you can just also use string(partsOfUrl)
   fmt.Println(anotherUrl)



}	