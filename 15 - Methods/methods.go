package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (user user) save() {
	fmt.Printf("Saving user data of %s on the database\n", user.name)
}

func (user user) majority() bool {
	if user.age >= 18 {
		return true
	}
	return false
}

func (user *user) happyBirthday() {
	user.age++
}

func main() {

	user1 := user{"User 1", 18}
	fmt.Println(user1)
	user1.save()
	fmt.Println(user1.majority())

	user2 := user{"User 2", 20}
	user2.save()
	fmt.Println(user2.majority())

	user2.happyBirthday()
	fmt.Println(user2.age)

}
