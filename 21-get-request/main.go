package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main() {
	fmt.Println("Get request tutorial")
	url:="http://localhost:5500/get"
	PerformGetRequest(url)

}

func PerformGetRequest(url string) {
 response, err := http.Get(url)
 if err!=nil {
	fmt.Println("Error Occured")	
 }
 defer response.Body.Close()
 fmt.Println("Status Code:", response.StatusCode)
 
 fmt.Println("ContentLength", response.ContentLength)
 content,_ := ioutil.ReadAll(response.Body)
 fmt.Println("Content:", string(content))
 var responseString strings.Builder//This response string is used to write the content to the string
 
 byteCount,_ := responseString.Write(content)//Now responseString contains the all content of the response
 fmt.Println("Byte Count:", byteCount)
 fmt.Println(responseString.String())
}	