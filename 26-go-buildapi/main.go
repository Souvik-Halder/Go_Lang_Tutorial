package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model of course -file
type Course struct{
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	CoursePrice int `json:"price"`	
	Author *Author `json:"author"`//This is the type of the author 

}

//fake DB

var courses []Course

// middleware , helper functions

func (c *Course)IsEmpty() bool{
	return   c.CourseName == "" 
}


type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
	
}

func main(){
	fmt.Println("This is the build api example")

	//creating routes
	r:= mux.NewRouter()
	
	//seeding - injecting some data in the DB

	courses=append(courses,Course{CourseId: "2",CourseName: "Reactjs",CoursePrice: 344,Author: &Author{Fullname: "Souvik",Website: "lco.dev"} })
	courses=append(courses,Course{CourseId: "4",CourseName: "MernStack",CoursePrice: 209,Author: &Author{Fullname: "Hitesh",Website: "lco.dev"} })

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")


	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")




	//listen to a port
	log.Fatal(http.ListenAndServe(":4000",r))
	
}

//controllers functions

//serve home route

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1> Welcome to the Go lang course</h1>"))
}

func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)//This will give the params in the url
	
	for _,item := range courses{
		if item.CourseId == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("No courses found with the id")
	return
}

func createOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Create one course")
	w.Header().Set("Content-Type","application/json")

	//what if : body is empty
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	
	//what about the body is - {}
	var course Course
  json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	
	
	for _,item:=range courses{
		
		if item.CourseName == course.CourseName{
			json.NewEncoder(w).Encode("You can't insert the data this course is already exist")
			return
		}
	}


	//generate unique id
	//append course into courses
	rand.Seed(time.Now().UnixNano())//for random number generation
	course.CourseId = strconv.Itoa(rand.Intn(1000000))//convert int to string
	courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
fmt.Println("update one course")
w.Header().Set("Content-Type","application/json")

params:= mux.Vars(r)

//loop ,id ,remove ,add with my ID

for index,course := range courses{
	if course.CourseId == params["id"]{
		courses = append(courses[:index],courses[index+1:]...)
		var course Course
		_ = json.NewDecoder(r.Body).Decode(&course)
		course.CourseId=params["id"]//overwriting the id 
		courses=append(courses,course)
		json.NewEncoder(w).Encode(course)//sending the course as a response json
		return
	}
}
//TODO: send a response when id is not found
}


func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("delete one course")
	w.Header().Set("Content-Type","application/json")
	
	params:= mux.Vars(r)//getting the params from the url

	for index,course:= range courses{
		if course.CourseId == params["id"]{
			courses= append(courses[:index],courses[index+1:]...)
			break;
		}
	}
 json.NewEncoder(w).Encode("The item has been delted")
 return
	
}