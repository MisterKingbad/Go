package main

import (
	"os"
	"fmt"
	"bufio"
)

// manipulação de arquivos

func main() {
	// criar arquivo
	f, err := os.Create("cmw.txt")
	if err != nil {
		panic(err)
	}

	// gravar dados no arquivo.
	stringLoren := "A expressão Lorem ipsum em design gráfico e editoração é um texto padrão em latim utilizado na produção gráfica para preencher os espaços de texto em publicações para testar e ajustar aspectos visuais antes de utilizar conteúdo real."

	// tam, err := f.WriteString(stringLoren)
	tam, err := f.Write([]byte(stringLoren))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado comsucesso! Tamanho: %d bytes\n", tam)
	
	f.Close()

	// Leendo um arquivo

	// file, err := os.Open("cmw.txt")
	file, err := os.ReadFile("cmw.txt")

	if err != nil {
		panic(err)
	}
	
	fmt.Println(string(file))

	// Leendo um arquivo de pouco em pouco

	arq, err := os.Open("cmw.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arq)

	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("cmw.txt")
	if err != nil {
		panic(err)
	}

}