package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementing i")
		//time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j <= 10; j += 2 {
		fmt.Println("Incrementing j", j)
		//time.Sleep(time.Second)
	}

	names := []string{"Name1", "Name2", "Name3"}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for value := range names {
		fmt.Println(value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

	for index, letter := range "Word" { // ASC values returned
		fmt.Println(index, letter)
	}

	for index, letter := range "Word" {
		fmt.Println(index, string(letter))
	}

	user := map[string]string{
		"name":     "Ivan",
		"lastName": "Koschelny",
	}

	for index, value := range user {
		fmt.Println(index, value)
	}

	for {
		//infinit loop
		break
	}

}
