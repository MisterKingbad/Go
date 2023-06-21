package main

import "fmt"

func main() {
	ev := []string{"test", "test2", "test3", "test4"}
	ev = append(ev[:0], ev[1:]...)
	// fmt.Println(ev1)
	fmt.Println(ev)
}