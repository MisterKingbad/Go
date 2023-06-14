package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int	`json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta) // Marshal guardar valor na variavel
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) // NewEncoder ja serializa jogango onde você quiser
	if err != nil {
		fmt.Println(err)
	}

	jsonPuro := []byte(`{"n":2, "saldo":200}`)
	var contaX Conta

	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(contaX.Saldo)

}
