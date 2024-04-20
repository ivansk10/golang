package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil { // if no panic on the application, it will be called but recover return would be nil
		fmt.Println("Recovered succeed")
	}
}

func approved(n1, n2 float64) bool {
	defer recoverExecution()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media is 6 and can't be") // stop what is running, it will call all defer functions
}

func main() {

	fmt.Println(approved(6, 6))
	fmt.Println("After execution")

}
