package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing of main archive")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("ivansaad23@gmail.com")
	erro2 := checkmail.ValidateFormat("123.com")
	fmt.Println(erro)
	fmt.Println(erro2)
}
