package main

import "fmt"

func weekDay(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
}

func weekDay2(number int) string {
	var weekDay string
	switch {
	case number == 1:
		weekDay = "Sunday"
		fallthrough
	case number == 2:
		weekDay = "Monday"
	case number == 3:
		weekDay = "Tuesday"
	case number == 4:
		weekDay = "Wednesday"
	case number == 5:
		weekDay = "Thursday"
	case number == 6:
		weekDay = "Friday"
	case number == 7:
		weekDay = "Saturday"
		fallthrough
	default:
		weekDay = "Invalid number"
	}
	return weekDay
}

func main() {

	day := weekDay(7)
	fmt.Println(day)
	day2 := weekDay(200)
	fmt.Println(day2)

	day3 := weekDay2(6)
	fmt.Println(day3)
	day4 := weekDay2(200)
	fmt.Println(day4)

	dayfallThrough := weekDay2(1)
	fmt.Println(dayfallThrough)
	dayfallThrough2 := weekDay2(7)
	fmt.Println(dayfallThrough2)

}
