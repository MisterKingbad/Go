package main

import (
	"fmt"
)

const a = "Holo Moto"
type ID int // criando um tipo

var (
	b bool    = false  
	c int     = 19
	d string  = "MisterKing"
	e float64 = 1.7
	f ID      = 1
)

func main () {
	var meuArray [3]int
	meuArray [0] = 1
	meuArray [1] = 2
	meuArray [2] = 50

	fmt.Println(3)
	fmt.Println(len(meuArray) -1) 
	fmt.Println(meuArray[len(meuArray) -3])

	for i, v := range meuArray { // O indice são as posições e o valor é o valor.
		fmt.Printf("O valor do indice é %d eo valor é %d\n", i, v)
	}
}