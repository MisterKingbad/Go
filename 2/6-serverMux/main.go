package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const Url = "https://viacep.com.br/ws/"

func main() {
	mux := http.NewServeMux() // com server mux temos mais controle do nosso servidor e podemos criar  varios endpoints
	mux.HandleFunc("/", HomeH)
	mux.Handle("/blog", blog{title: "Meu blog"})
	mux.Handle("/buscar/", Endereco{})
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

func (e Endereco) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusNotFound) // resposta do cabeçalho
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

//funcoes para fazer a requisicao

func BuscaCep(cep string) (*Endereco, error) {
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
