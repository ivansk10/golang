package main

import "fmt"

func main() {
	sum := 1 + 2
	sub := 1 - 3
	div := 10 / 4
	multi := 10 * 5
	rest := 10 % 2

	fmt.Println(sum, sub, div, multi, rest)

	//var numero1 int16 = 10
	//var numero2 int32 = 25
	//Error due to data type
	//soma := numero1 + numero2
	//fmt.Println(soma)

	var num1 int16 = 10
	var num2 int16 = 25
	sum2 := num1 + num2
	fmt.Println(sum2)
	fmt.Println("------------------")
	//Fim aritmeticos

	//Atribuição

	var variable1 string = "String"
	variable2 := "String 2"
	fmt.Println(variable1, variable2)
	fmt.Println("------------------")
	//End Atribuicao

	//Operadores relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	fmt.Println("------------------")
	// Fim operadores relacionais

	//Operadores logicos
	trueBool, falseBool := true, false
	fmt.Println(trueBool && falseBool)
	fmt.Println(trueBool || falseBool)
	fmt.Println(!trueBool)
	fmt.Println(!falseBool)
	fmt.Println("------------------")

	// Fim operadores logicos

	//Operadores unarios
	num := 10
	num++
	num += 15
	num--
	num -= 15
	num *= 3
	num /= 10
	num %= 3
	fmt.Println(num)
	fmt.Println("------------------")

	//Fim dos operadores unarios

}
