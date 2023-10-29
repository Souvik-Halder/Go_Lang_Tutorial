package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Post request tutorial")

	url:="http://localhost:5500/post"
	PerformPostRequest(url)

}

func PerformPostRequest(url string){
	
	requestBody := strings.NewReader(`
	{
		"name":"Souvik",
		"course":"Golang",
		"blog":"https://www.golangprograms.com",
		"price":0,
		"platform":"Golangprograms.com"
	}`)

	response,err := http.Post(url,"application/json",requestBody)
	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()//alsways close the response body
	// var responseString strings.Builder
	content,_:= ioutil.ReadAll(response.Body)
	fmt.Println("Content:", string(content))
	// byteCount,_:= responseString.Write(content)
	// fmt.Println("Byte Count:", byteCount)
	// fmt.Println(responseString.String())

}	
