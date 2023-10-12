package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on maps")

	languages := make(map[string]string, 5) //This is how we can create a map in go language the first string is the key and the second string is the value
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["CS"] = "Csharp"
	languages["TS"] = "Typescript"
	delete(languages, "RB") //This is how we can delete the element from the map
	fmt.Println(languages)
	fmt.Println(languages["PY"]) //This will print the value of the py key

	//looping through the map
	for key, value := range languages {
		fmt.Printf("For key %v the value is %v\n", key, value)
	}
}
