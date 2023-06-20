package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id int64
	Msg string
}

func main() { 
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Rabb"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i,"Kaf"}
			c2 <- msg
		}
	}()

	// for i:= 0; i < 2; i++ {
		for {

		select {
		case msg := <-c1:
			fmt.Printf("Received from Rabb: ID: %d - %s\n", msg.id, msg.Msg)
		case msg := <-c2:
			fmt.Printf("Received from Kaf: ID: %d - %s\n", msg.id, msg.Msg)
		case <-time.After(time.Second * 3):
			println("timeout")
	
		// default:
		// 	println("default")
		}
	}

}