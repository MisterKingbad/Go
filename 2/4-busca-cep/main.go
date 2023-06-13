package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/"+ cep +"/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		fmt.Println(cep)

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		var data Endereco
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		fmt.Println(data)

		file, err := os.Create("cidade.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.Write([]byte(fmt.Sprintf("cep:%s,\nlocalidade:%s,\nuf:%s\n", data.Cep, data.Localidade, data.Uf)))
		fmt.Println("Arquivo criado com sucesso!")
	}
}
