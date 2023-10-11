package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our pizza 1 to 5: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')                                //here the input consists the number and the \n also with it if you enter the 4 in input and press enter so the input will be 4\n because u press the enter after giving the input so we need to remove it by strings.TrimSpace(input) function
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // 64 is the bit size // TrimSpace removes the \n
	fmt.Println("Thanks for rating, ", numRating)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating, ", numRating+1)
	}

}
