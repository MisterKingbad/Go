package main

func main() {
	ch := make(chan string, 2)	
	ch <- "Holo"
	ch <- "world war"

	println(<-ch)
	println(<-ch)
}