package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Welcome to modules in golang");
	Greeter();
	r:= mux.NewRouter();
	r.HandleFunc("/",serveHome).Methods("GET");
	log.Fatal(http.ListenAndServe(":4000",r))
}

func Greeter(){
	fmt.Println("Hello from Greeter");
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome to my awesome site</h1>"))
}