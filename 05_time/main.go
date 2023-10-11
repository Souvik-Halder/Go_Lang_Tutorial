package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to our app")
	parseTime := time.Now()
	fmt.Println(parseTime)
	fmt.Println(parseTime.Format("02-01-2006 15:04:05 Monday")) //This is standart format of the date
	createDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	//formated version of createDate
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
