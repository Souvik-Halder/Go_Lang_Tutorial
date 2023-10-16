package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on loops")

	days:= []string{"Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"}
	fmt.Println(days);
	//for loop
	// for d:=0;d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	for index,day :=range days{
		fmt.Println("index is ",index,"value is ",day)
	}

	roughValue :=0
	for roughValue<10{
		// if roughValue == 2{
		// 	goto lco
		// }//This used to goto to the label 
		if roughValue ==4{
			continue //This wil skip the 4 value
		} // This is used to continue the loop if 4 occurs so it basically will skip the value 4
		if roughValue ==5 {
			break; // This will break after the 5 value
		} //This will break the loop at 5
		fmt.Println("Value is",roughValue)
		roughValue++;
	}

// lco:
// 	fmt.Println("This is a label using goto u can jump in that")
}