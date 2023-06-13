package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	numeros := []int{1, 2 , 3, 4, 5}

	for p, n := range numeros {
		fmt.Println(p, n)
	}

	frutas := []string{"banana", "laranja", "uva", "maçã", "abacate", "abacaxi"}

	for p, f:= range frutas {
		fmt.Printf("A %v esta na posição %v\n", f, p)
	}

	i :=0

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// for {
	// 	fmt.Println("HELLO WORLD!") // LOOP INIFNITO
	// }
}
