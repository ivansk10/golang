package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int8 = 100
	fmt.Println(numero)
	//var numero2 int16 = 10000000000000000
	//fmt.Println(numero2)
	//var numero3 int32 = 10000000000000000 Error due to int capacity
	//.Println(numero3)
	var numero4 int64 = 10000000000000000
	fmt.Println(numero4)
	var numero5 int = 100000000000000
	fmt.Println(numero5)
	numero6 := 100000000000
	fmt.Println(numero6)
	numero7 := -100000000000
	fmt.Println(numero7)
	var numero8 uint32 = 100
	fmt.Println(numero8)
	//int32 = rune
	var numero9 rune = 12345
	fmt.Println(numero9)
	//byte = uint8
	var numero10 byte = 123
	fmt.Println(numero10)

	//Numeros Reais

	var numeroReal11 float32 = 10.000000000000000001
	fmt.Println(numeroReal11)
	var numeroReal12 float64 = 1000000000000000.0
	fmt.Println(numeroReal12)
	//Fim Numeros reais

	//String

	var str string = "Text"
	fmt.Println(str)
	var str2 string = "Text2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	//Fim Strings

	//Valores iniciais

	var texto string
	fmt.Println(texto)
	var num int16
	fmt.Println(num)

	//Fim

	//Boolean

	var booleano bool = true
	var booleano2 bool = false
	var booleano3 bool
	fmt.Println(booleano)
	fmt.Println(booleano2)
	fmt.Println(booleano3)
	//Fim Boolean

	//Error

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Intern error")
	fmt.Println(erro2)

}
