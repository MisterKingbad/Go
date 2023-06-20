package main

import "fmt"

func receba(nome string, hello chan<- string) { // se a seta estiver do lado direito quer dizer q ele esta recebendo dados
	hello <- nome
}

func ler(data <-chan string) { // se a seta estiver do lado esquerdo quer dizer q ele esta entregando os dados/ resultados
	fmt.Println(<-data)
}

//Thread 1
func main() {
	hello := make(chan string)
	go receba("Irineu", hello)
	ler(hello)
}
