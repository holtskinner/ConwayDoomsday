package main

import (
	"fmt"
	"time"
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

var doomsdays = [12]int{3, 28, 14, 4, 9, 6, 11, 8, 5, 10, 7, 12}

func main() {

	// var month, day, year int

	// for month < 1 || month > 12 {
	// 	fmt.Printf("Enter the Month: ")
	// 	fmt.Scanln(&month)
	// }

	// // TODO: More error checking
	// for day < 1 || day > 31 {
	// 	fmt.Printf("Enter the Day: ")
	// 	fmt.Scanln(&day)
	// }

	// // Gregorian Calendar released in 1582, 3000 arbitrary
	// for year < 1582 || year > 2700 {
	// 	fmt.Printf("Enter the year: ")
	// 	fmt.Scanln(&year)
	// }

	start := time.Now()

	for i := 0; i < 1000; i++ {
		for y := 1582; y < 2700; y++ {
			for m := 1; m <= 12; m++ {
				for d := 1; d <= 28; d++ {

					doomsday(m, d, y)
				}
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Println(elapsed)

}

// Mathematical Equivelent
func fastDoomsday(month, day, year int) Weekday {

	d := 2 + year + (year / 4) - (year / 100) + (year / 400)
	// leapyear
	var l int

	if (month == 1 || month == 2) && (year%4 == 0 && year%400 != 0) {
		l = 1
	}

	h := (day - doomsdays[month-1] - l + d) % 7

	return Weekday(h)

}

// Conway's Original Doomsday Algorithm
func doomsday(month, day, year int) Weekday {

	var i int

	// Find anchor day for century
	switch year / 100 {
	case 15, 19, 23:
		i = 3
	case 16, 20, 24:
		i = 2
	case 17, 21, 25:
		i = 0
	case 18, 22, 26:
		i = 5
	}

	// Find doomsday for year

	// Last two digits of year
	y := year % 100

	a := y / 12

	b := y % 12

	c := b / 4

	// Day of the week for the doomsday
	d := a + b + c + i

	// leapyear
	var l int

	if (month == 1 || month == 2) && (year%4 == 0 && year%400 != 0) {
		l = 1
	}

	h := (day - doomsdays[month-1] - l + d) % 7

	return Weekday(h)
}
