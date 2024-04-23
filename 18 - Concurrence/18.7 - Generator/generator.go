package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("Hello World")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string { // <-chan only to receive values
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
