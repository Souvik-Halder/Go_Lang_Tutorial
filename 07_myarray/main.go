package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to a class on arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Orange"
	fruitList[3] = "Peach"
	fmt.Println("FruitList is: ", fruitList)
	fmt.Println("The length of the FruitList is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "onion"}
	sort.Strings(vegList[:])
	fmt.Println("VegList is: ", vegList)

}
