package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// w.Write([]byte("Hello World!"))
	log.Println("Request iniciada")
	defer log.Println("Request Finalizada")

	select {
	case <-time.After(5* time.Second):
		log.Println("Requisição processada com sucesso")
		w.Write([]byte("Requisição processada com sucesso"))

	case <- ctx.Done():
		log.Println("Requisição cancelada pelo cliente")
		http.Error(w, "Requisição cancelada pelo cliente", http.StatusRequestTimeout)
	}
}