package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := multiplexar(write("Hello World"), write("Hello World 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}

}

func multiplexar(channel1, channel2 <-chan string) <-chan string {
	exitChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-channel1:
				exitChannel <- message
			case message := <-channel2:
				exitChannel <- message
			}
		}
	}()
	return exitChannel
}

func write(text string) <-chan string { // <-chan only to receive values
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
