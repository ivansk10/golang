package main

import "fmt"

func sum(numbers ...int) (total int) {
	total = 0
	for _, value := range numbers {
		total += value
	}
	return
}

func write(text string, numbers ...int) { //this parameter sequence must be followed
	for _, value := range numbers {
		fmt.Println(text, value)
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 200, 500))
	write("Hello World", 10, 20, 30, 40, 50)
}
