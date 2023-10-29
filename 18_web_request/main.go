package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url="https://lco.dev"

func main(){
	fmt.Println("Web request")
     response,err := http.Get(url)
	 if (err!=nil){
		 panic(err)
	 }
	 fmt.Printf("Response is of type: %T\n",response);
	

	defer response.Body.Close()//it is caller responsibility to close the response body

	 databytes,err := ioutil.ReadAll(response.Body)	
	 if (err!=nil){
		 panic(err)
	 }
	 content:=string(databytes)
	 fmt.Println(content)
}