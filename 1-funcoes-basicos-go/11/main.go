package main

import "fmt"

type Cliente struct { // como se fosse o serializer django ou schema ninja ou interface nextjs
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	irineu := Cliente {
		Nome: "Irineu",
		Idade: 40,
		Ativo: true,
	}
	irineu.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo %t", irineu.Nome, irineu.Idade, irineu.Ativo)
}
