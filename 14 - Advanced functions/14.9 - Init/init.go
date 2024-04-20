package main

import "fmt"

var n int

func init() { // different from main, we can have one per file
	fmt.Println("Init being executed before main()")
	n = 10
}

func main() {
	fmt.Println("main() being executed")
	fmt.Println(n)
}
