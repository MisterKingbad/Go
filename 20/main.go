package main

import (
	"fmt"
	"curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Printf("O resultado: %v", s)
	fmt.Printf(matematica.A)
	

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf(carro.Marca)

	fmt.Println(carro.Andar())
}