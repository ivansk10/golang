package main

import "fmt"

func mathCalc(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	sum, sub := mathCalc(10, 20)
	fmt.Println(sum, sub)

}
