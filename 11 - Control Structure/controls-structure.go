package main

import "fmt"

func main() {
	number := 10

	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Smaller than 15")
	}

	if anotherNumber := number; anotherNumber > 10 {
		fmt.Println("Greater than 0", anotherNumber) //anotherNumber can't be accessed outside of this scope
	} else if anotherNumber3 := 1; anotherNumber3 < 15 {
		fmt.Println("Smaller than 15")
	}
}
