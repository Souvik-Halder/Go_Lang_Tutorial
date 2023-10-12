package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to a class on slices")
	var fruitList = []string{"Appele", "Tomato", "Orange"} //This is just like dynamic array we don't need to specify the size of the  array in the slice
	fruitList = append(fruitList, "Peach", "Banana")       //way to add element in the slice
	fmt.Println("FruitList is: ", fruitList)
	fmt.Printf("Type of the is:%T\n", fruitList)

	fruitList = append(fruitList[1:3]) //This is how we can delete the element from the slice
	fmt.Println("FruitList is: ", fruitList)

	highScores := make([]int, 4) //By this way you will create a slice of size 4
	highScores[0] = 234          //This is not the new memory allocation whenever the append method is called the go allocates the memory
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 78// u can't use this to add element in the slice
	highScores = append(highScores, 555, 666, 777) //This is how we can add element in the slice
	//because here the size is initialized to 4 and we are trying to add 5th element in the slice
	//so it will create a new array of size 8 and copy the previous element in the new array and then add the new element
	sort.Ints(highScores)
	fmt.Println("HighScores is: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //This will return true if the slice is sorted

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //This is how we can delete the element from the slice This is the way to delete the specific index  element from the slice
	fmt.Println(courses)
}
