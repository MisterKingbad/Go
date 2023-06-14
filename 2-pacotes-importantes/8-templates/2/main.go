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
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := t.Execute(os.Stdout, curso) // stdout imprime na tela os dados e o curso esta passando as variaveis dentro dele para utilizar
	if err != nil {
		panic(err)
	}
}
