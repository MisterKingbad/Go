package main

import "fmt"

func main() {
	var minhaVar interface{} = "iri"

	fmt.Println(minhaVar)
	println(minhaVar.(string)) // Reafirmando o tipo da variavel

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de 'res' é %v e o resultado de 'ok' é %v\n", res, ok)
	res1, _ := minhaVar.(string)
	fmt.Printf("O valor de 'res' é %v\n", res1)
}