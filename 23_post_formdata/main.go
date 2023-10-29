package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main(){
	fmt.Println("Post form data tutorial")
	url:="http://localhost:5500/postform"
	PerformPostFormRequest(url)

}

func PerformPostFormRequest(myurl string){

	//formData
	data:= url.Values{}
	data.Add("name","Souvik")
	data.Add("course","Golang")
	data.Add("blog","https://www.golangprograms.com")
	data.Add("price","0")
	data.Add("platform","Golangprograms.com")

	fmt.Println(data)

	response,err := http.PostForm(myurl,data)
	if(err!=nil){
		panic(err)
	}
	defer response.Body.Close()
	content,_ := ioutil.ReadAll(response.Body)
	fmt.Println("Content:", string(content))

}