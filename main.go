package main

import (
	"fmt"
)

// Weekday Day of the Week (int)
type Weekday int

// Days of the week
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[w]
}

func main() {

	var month, day, year int

	for month < 1 || month > 12 {
		fmt.Printf("Enter the Month: ")
		fmt.Scanln(&month)
	}

	// TODO: More error checking
	for day < 1 || day > 31 {
		fmt.Printf("Enter the Day: ")
		fmt.Scanln(&day)
	}

	// Gregorian Calendar released in 1582, 3000 arbitrary
	for year < 1582 || year > 3000 {
		fmt.Printf("Enter the year: ")
		fmt.Scanln(&year)
	}

	fmt.Println(month, day, year)

}

func doomsday(month, day, year int) Weekday {
	return 1
}
