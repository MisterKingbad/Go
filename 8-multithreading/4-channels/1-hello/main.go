package main

import "fmt"

//Thread 1
func main() {
	canal := make(chan string) // "chan" channel e Vazio

	// Thread 2
	go func() {
		canal <- "Hello new world!" // prenchendo o canal com a string
	}()

	// Thread 3
	msg := <-canal // canal esvaziando
	fmt.Println(msg)
}