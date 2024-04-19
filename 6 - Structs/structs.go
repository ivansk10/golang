package main

import "fmt"

type user struct {
	name string
	age  uint8
	address
}

type address struct {
	publicPlace string
	number      uint8
}

func main() {

	var user1 user
	user1.name = "Ivan Koschelny"
	user1.age = 28
	fmt.Println(user1)

	adress1 := address{"Rua da Paz", 30}

	user2 := user{"Ivan Koschelny", 28, adress1}
	fmt.Println(user2)

	user3 := user{age: 28, address: address{"Rua da paz", 25}}
	number1 := user3.number
	fmt.Println(number1)
	fmt.Println(user3)

}
