package main
import "fmt"

func main() {
	salarios := map[string]int{
		"Iri": 1000,
		"Neu": 2850,
		"spooderMan": 10000,
	}

	for nome, salario := range salarios {
		fmt.Printf("O %v ganha R$%v por mês!\n", nome, salario)
	}

	for _, salario := range salarios { // "_" ignora o outro valor para nao ser usado
		fmt.Printf("Valores sendo pago para os melhores funcionarios $%v por mês\n",salario)
	}

	fmt.Println(salarios["Iri"])
	delete(salarios, "Iri")
	salarios["ChingLing"] = 15000
	fmt.Println(salarios["ChingLing"])

	sal := make(map[string]int)
	fmt.Println(sal)
	sal1 := map[string]int{}
	fmt.Println(sal1)
}