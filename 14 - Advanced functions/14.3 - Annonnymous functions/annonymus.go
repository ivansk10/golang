package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello World")
	}()

	func(text string) {
		fmt.Println(text)
	}("Parameters to annonymus function")

	functionReturn := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text) // %s for string, %d for number
	}("Parameters to annonymus function")

	fmt.Println(functionReturn)

}
