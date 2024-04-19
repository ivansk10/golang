package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	variavel2++
	fmt.Println(variavel1, variavel2)

	//Pointer is a memory reference

	var variavel3 int = 100
	var pointer *int = &variavel3
	fmt.Println(variavel3, pointer)
	variavel3++
	fmt.Println(variavel3, *pointer) // unreference

}
