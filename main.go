package main

import (
	"fmt"
)

var weekdays = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func main() {

	var month, day, year int

	fmt.Printf("Enter the Month: ")
	fmt.Scanln(&month)

	fmt.Printf("Enter the Day: ")
	fmt.Scanln(&day)

	fmt.Printf("Enter the year: ")
	fmt.Scanln(&year)

	fmt.Println(month, day, year)
}

func doomsday() {

}
