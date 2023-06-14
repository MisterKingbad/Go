package main

import (
	"net/http"
	// "io/ioutil"
	"bytes"
	"io"
	"os"
)



func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"nome": "iri"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(body))
}