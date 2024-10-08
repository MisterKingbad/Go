package main

import (
	"fmt"
	"time"
)

func worker(wokerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", wokerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 1000000

	for i := 0; i< QtdWorkers; i++ {
		go worker(i, data)
	}
	
	for i := 0; i < 10000000; i++ {
		data <- i
	}
}