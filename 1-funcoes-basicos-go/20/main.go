package main

import (
	"fmt"
	"curso-go/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Printf("O resultado: %v", s)
	fmt.Printf(matematica.A)
	

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf(carro.Marca)

	fmt.Println(carro.Andar())
	fmt.Println(uuid.New())
}