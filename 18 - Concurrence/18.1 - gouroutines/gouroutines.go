package main

import (
	"fmt"
	"time"
)

func main() {
	//Concurrence != parallelism
	go write("Hello World") //goroutine
	write("Programming in go")

}

func write(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
