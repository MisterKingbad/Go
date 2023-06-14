package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Golang", 200}
	tmp := template.New("CursoTemplate")

	tmp, err := tmp.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, curso) // stdout imprime na tela os dados e o curso esta passando as variaveis dentro dele para utilizar
	if err != nil {
		panic(err)
	}
}
