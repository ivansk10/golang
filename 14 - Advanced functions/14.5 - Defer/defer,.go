package main

import "fmt"

func function1() {
	fmt.Println("Executiting Function 1")
}

func function2() {
	fmt.Println("Executiting Function 2")
}

func approved(n1, n2 float32) bool {
	defer fmt.Println("Calculated. Result will be returned")
	fmt.Println("Is student approved?")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	function1()
	function2()

	fmt.Println("-----------")

	defer function1() // defer function until last gasp
	function2()
	fmt.Println("-----------")
	fmt.Println(approved(7, 8))

}
