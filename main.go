package main

import "fmt"

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

var doomsdays = [12]int{3, 28, 14, 4, 9, 6, 11, 8, 5, 10, 7, 12}

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
	for year < 1582 || year > 2700 {
		fmt.Printf("Enter the year: ")
		fmt.Scanln(&year)
	}

}

// Mathematical Equivelent
func doomsday2(month, day, year int) Weekday {

	d := 2 + year + (year / 4) - (year / 100) + (year / 400)
	// leapyear
	var l int

	if (month == 1 || month == 2) && (year%4 == 0 && year%400 != 0) {
		l = 1
	}

	h := (day - doomsdays[month-1] - l + d) % 7

	return Weekday(h)

}
