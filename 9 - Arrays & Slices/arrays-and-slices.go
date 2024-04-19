package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "Position 1"
	array2[1] = "Position 2"
	fmt.Println(array2)

	array3 := [3]string{"Position 1", "Position 2", "Position 3"}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)
	//array4[6] = 5 Out of bounds

	fmt.Println("------- SLICES -----")

	slice := []int{10, 11, 12, 13}
	fmt.Println(slice)
	slice = append(slice, 14)
	fmt.Println(slice)

	slice2 := array3[1:2] // the slice of the array does not consider last number position.
	fmt.Println(slice2)
	array3[1] = "Position changed"
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	fmt.Println("----- Intern arrays------")

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 11)
	fmt.Println(slice3)
	slice3 = append(slice3, 12)
	fmt.Println(slice3)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // Every time you pass the max capacity it will double in 2 times the capacity based on the number you exceed.

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
