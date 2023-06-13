package main

import (
	// "io"
	// "fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

const Url = "https://viacep.com.br/ws/"

func main() {
	http.HandleFunc("/api/buscar/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}


func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/buscar/" {
		w.WriteHeader(http.StatusFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cep)
}


func BuscaCep(cep string) (*Endereco, error){
	req, err := http.Get(Url + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}
	var c Endereco 
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil

}
