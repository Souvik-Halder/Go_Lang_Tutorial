package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`//This is alices which means when the Json is made by that 
	Price int    `json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"`//This will not send the password in the json data
	Tags []string  `json:tags,omitempty`//This will not show the tags if it is empty
}

func main(){
	fmt.Println("Welcome to Json data tutorial")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson(){
	lcoCourses := []course{
		{"Reactjs",30,"hi.com","abc123",[]string{"web-dev","js"}},
		{"Mongodb",20,"hi.com","abc123",[]string{"web-dev","js"}},
		{"Golang",0,"hi.com","abc123",[]string{"web-dev","js"}},
		{"Python",0,"hi.com","abc123",[]string{"web-dev","js"}},

	}

	finalJson,err := json.MarshalIndent(lcoCourses,"","\t");//This will structure the JSON data in a proper way by indentation
	if( err != nil){
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
	
}

func DecodeJson(){
	jsonDatafromWeb := []byte(`
	{
		"coursename": "Reactjs",
		"price": 30,
		"website": "hi.com",
		"Tags": ["web-dev","js"]
}`)
 
    var lcoCourses course //This is the struct which will hold the data

	//cheching the validity of the JSON
	checkValid := json.Valid(jsonDatafromWeb)//This will check if the json data is valid or not
	
	if checkValid{
		fmt.Println("Json is valid")
		//This is the Unmarshal function exactly opposite of the Marshal() function will conver json data to struct opposite of the Marshal() function
		json.Unmarshal(jsonDatafromWeb,&lcoCourses)//This will convert the json data into the struct
	
		fmt.Println(lcoCourses.Name,lcoCourses.Platform,lcoCourses.Tags)
	}else{

		fmt.Println("Json is not valid")
	}


	//some cases where you want to ge the data in the key value pair

	var myOnlineData map[string] interface{}//This is a map which will hold the data
	json.Unmarshal(jsonDatafromWeb,&myOnlineData)//This will convert the json data into the map
	fmt.Printf("%v\n",myOnlineData)

	for k,v := range myOnlineData{
		fmt.Println("Key is:",k,"Value is:",v)
	}
	
}