package main

import "fmt"

func inverteSignal(number int) int {
	return number * -1
}

func invertedSignalWithPointer(number *int) { // does not need return because we are chaging the value inside of the box
	*number = *number * -1
}

func main() {
	number := 20
	invertedSignal := inverteSignal(number)
	fmt.Println(invertedSignal)
	fmt.Println(number)

	newNumber := 40
	fmt.Println(newNumber)
	invertedSignalWithPointer(&newNumber)
	fmt.Println(newNumber)

}
