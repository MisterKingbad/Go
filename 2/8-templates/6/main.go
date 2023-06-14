package main

import (
	"net/http"
	"html/template"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// t := template.Must(template.New("content.html").ParseFiles(templates...))
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))
		err := t.Execute(w, Cursos{
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
	})
	http.ListenAndServe(":8080", nil)
}
