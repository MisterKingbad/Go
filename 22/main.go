package main

import "fmt"

func main() {
	a:= 1
	b:= 2
	c:= 3

	numeros :=[]int{1, 2, 3}

	for _, v:= range numeros {
		switch v {
		case 1:
			fmt.Println("a")
		case 2:
			fmt.Println("b")
		case 3:
			fmt.Println("c")
		default:
			fmt.Println("d")
		}
	}


	if a > b  && c > a{
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
