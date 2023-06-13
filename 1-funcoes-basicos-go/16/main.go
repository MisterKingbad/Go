package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	fmt.Printf("Saldo simulado total %v\n",c.saldo)
	return c.saldo
}

func NewConta() *Conta {
	return &Conta{
		saldo: 0,
	}
}

func (c Cliente) andou() {
	c.nome = "Irineu"
	fmt.Printf("O clinte %v andou\n", c.nome)
}

func main() {
	iri := Cliente{
		nome: "Iri",
	}
	iri.andou()
	fmt.Printf("O valor da struct com nome %v\n", iri.nome)

	conta := Conta{
		saldo: 1000,
	}
	conta.simular(2000)
	fmt.Printf("Saldo atual %v\n",conta.saldo)
}
