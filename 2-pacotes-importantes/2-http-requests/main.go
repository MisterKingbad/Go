package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close() // fechar para não vazar e o defer executa por ultimo depois ou por ultimo

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

	

}
