package main

import "fmt"

func main() {
	user := map[string]string{
		"name":     "Ivan",
		"lastName": "Koschelny",
	}

	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["lastName"])

	user2 := map[string]map[string]string{
		"name": {
			"first":  "Ivan",
			"second": "Ivan2",
		},
		"curso": {
			"name":   "System Analysis",
			"campus": "PucPR",
		},
	}

	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)
	user2["age"] = map[string]string{
		"ageAtBeging": "10",
	}
	fmt.Println(user2)
}
