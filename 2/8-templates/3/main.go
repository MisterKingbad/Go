package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 500},
		{"Python", 100},
		{"Next", 50},
		{"React", 60},
		{"Django", 90},
		{"DRF", 40},
		{"Django Ninja", 40},
		{"html", 40},
		{"Docker", 40},
		{"Redis", 40},
		{"Apache", 40},
		{"Postgresql", 40},
		{"Mysql", 40},
		
	}) // stdout imprime na tela os dados e o curso esta passando as variaveis dentro dele para utilizar
	if err != nil {
		panic(err)
	}
}
