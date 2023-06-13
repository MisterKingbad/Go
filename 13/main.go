package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface { // somente usar metodos nao props
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct { // como se fosse o serializer django ou schema ninja ou interface nextjs
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O Cliente %s foi desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	irineu := Cliente{
		Nome:  "Irineu",
		Idade: 40,
		Ativo: true,
	}
	Desativacao(irineu)
	irineu.Desativar()
	irineu.Endereco.Logradouro = "Av Alberto"
	irineu.Endereco.Numero = 226
	irineu.Endereco.Cidade = "São Paulo"
	irineu.Endereco.Estado = "SP"

	fmt.Printf("Nome: %s, Idade: %d, Ativo %t, Logradouro: %s, Numero: %d, Cidade: %s, Estado: %s\n", irineu.Nome, irineu.Idade, irineu.Ativo, irineu.Endereco.Logradouro, irineu.Endereco.Numero, irineu.Endereco.Cidade, irineu.Endereco.Estado)

	fmt.Printf("Nome: %s, Idade: %d, Ativo %t, Logradouro: %s, Numero: %d, Cidade: %s, Estado: %s\n", irineu.Nome, irineu.Idade, irineu.Ativo, irineu.Endereco.Logradouro, irineu.Endereco.Numero, irineu.Endereco.Cidade, irineu.Endereco.Estado)

	minhaEmpresa := Empresa {
		Nome: "Nike clandestino",
	}
	Desativacao(minhaEmpresa)
}
