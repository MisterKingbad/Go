package main

import (
	"context"
	"net/http"
	"time"
	"os"
	"io"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil) // req com context
	if err  != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req) //executa a req 
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // fechar a conexao res
	io.Copy(os.Stdout, res.Body)
}