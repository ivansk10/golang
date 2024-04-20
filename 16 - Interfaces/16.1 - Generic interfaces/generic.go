package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(10)
	generic(true)

	fmt.Println(1, 2, "String", false, true, float64(1234)) // fmt.Println is a example of generic functions

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(212),
	}

	fmt.Println(mapa)
}
