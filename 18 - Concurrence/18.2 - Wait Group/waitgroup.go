package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		write("Programming in go 1")
		waitGroup.Done() // -1
	}()

	go func() {
		write("Programming in go 2")
		waitGroup.Done() //-1
	}()

	go func() {
		write("Programming in go 3")
		waitGroup.Done() //-1
	}()

	go func() {
		write("Programming in go 4")
		waitGroup.Done() //-1
	}()

	waitGroup.Wait() // 0

}

func write(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
