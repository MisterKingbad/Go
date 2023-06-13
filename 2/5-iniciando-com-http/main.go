package main

import (
	// "io"
	// "fmt"
	"net/http"
	// "encoding/json"
)

func main() {
	http.ListenAndServe(":8080", nil)
}