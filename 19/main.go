package main

import "fmt"

//Generics.


func som[T int | float64] (m map[string]T) T{
	var soma T 
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNum int

type Num interface {
	~int | ~float64
}

func som1[T Num] (m map[string]T) T{
	var soma T 
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Num](a T, b T) bool {
	if a > b {
		return true
	}
	return false
}

// func somInt(m map[string]int) int{
// 	var soma int 
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func somFloat(m map[string]float64) float64 {
// 	var soma float64
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

func main() {
	m := map[string]int{"Iri": 1000, "Neu": 2000, "Marilene": 3000}
	m1 := map[string]float64{"Iri": 1000.10, "Neu": 2000.10, "Marilene": 3000.10}

	m2 := map[string]MyNum{"Iri": 1000, "Neu": 2000, "Marilene": 3000}
	
	fmt.Println(som(m))
	fmt.Println(som(m1))
	fmt.Println(som1(m2))

	fmt.Println(Compara(10, 10))
}