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
	fmt.Printf("O tipo de E é %T ", f) // '%T' e para ver o tipo do variavel
	fmt.Printf("O tipo de E é %v", e) // '%v' e para ver o valor da variavel
}