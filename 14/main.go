package main

import "fmt"

func main() {
	// O ponteiro é um endereco na memoria que tem um valor as variaveis apontam para o endereco na memoria que tem um valor.

	a := 10 // Variavel principal

	//fmt.Println(&a) // Mostra onde ta guardado a variavel "a" com o valor 10.
	var ponteiro *int = &a // Mudar diretamente o dados da variavel "a" pelo endereço. O "*" quer dizer o endereco a rota que o valor esta.
	//fmt.Println(ponteiro)
	*ponteiro = 20 // Usando a variavel ponteiro que esta apontado para o endereco do "a" para mudar diretamento o valor de lá sem usar a variavel "a".
	b := &a
	fmt.Println(*b) // Quando tem "*" na frente da varivel e a variavel é um ponteiro ou tem um endereco nele, ele mostra o valor que tem nesse endereço.

	*b = 30
	fmt.Println(a)

}
