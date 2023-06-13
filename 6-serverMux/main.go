package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeH)
	mux.Handle("/blog", blog{title: "Meu blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeH(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem vindo!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

type cep struct {}
func (c cep) ServeHTTP(w http.ResponseWriter, r *http.Request)
