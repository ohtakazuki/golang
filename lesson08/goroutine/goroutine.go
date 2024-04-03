package main

import (
	"fmt"
	"time"
)

func printGoroutine() {
	for i := 0; i < 20; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go printGoroutine()

	for i := 0; i < 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(time.Second)
	}
}
