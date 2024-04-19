package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      uint8
	height   float32
}

type student struct {
	person
	course  string
	college string
}

func main() {
	person1 := person{"Ivan", "Koschelny", 28, 1.80}
	fmt.Println(person1)

	student1 := student{person1, "System analysis", "PucPR"}
	fmt.Println(student1)
	fmt.Println(student1.name)
	fmt.Println(student1.lastName)
	fmt.Println(student1.age)
	fmt.Println(student1.height)
}
