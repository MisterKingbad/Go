package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct { // como se fosse o serializer django ou schema ninja ou interface nextjs
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	irineu := Cliente{
		Nome:  "Irineu",
		Idade: 40,
		Ativo: true,
	}
	irineu.Ativo = false
	irineu.Endereco.Logradouro = "Av Alberto"
	irineu.Endereco.Numero = 226
	irineu.Endereco.Cidade = "São Paulo"
	irineu.Endereco.Estado = "SP"

	fmt.Printf("Nome: %s, Idade: %d, Ativo %t, Logradouro: %s, Numero: %d, Cidade: %s, Estado: %s", irineu.Nome, irineu.Idade, irineu.Ativo, irineu.Endereco.Logradouro, irineu.Endereco.Numero, irineu.Endereco.Cidade, irineu.Endereco.Estado)
}
