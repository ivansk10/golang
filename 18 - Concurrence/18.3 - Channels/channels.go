package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go write("Programming in go 1", channel)
	go write("Programming in go 2", channel)

	fmt.Println("After write execution")

	for message := range channel {
		fmt.Println(message)
	}

}

func write(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto
		time.Sleep(time.Second)
	}
	close(channel)
}
