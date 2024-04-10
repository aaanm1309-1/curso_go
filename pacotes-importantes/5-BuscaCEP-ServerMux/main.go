package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ViaCep struct {
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

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", BuscaCepHandler)
	log.Fatal(http.ListenAndServe(":8085", mux))

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", HomeHandler)
	mux2.Handle("/blog", blog{title: "My Blog"})
	log.Fatal(http.ListenAndServe(":8086", mux2))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Minha home do server 2"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" || len(cepParam) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cepRetorno, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cepRetorno)
	// ou
	// result, err := json.Marshal(cepRetorno)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func BuscaCep(cep string) (*ViaCep, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var returnViaCep ViaCep
	err = json.Unmarshal(body, &returnViaCep)
	if err != nil {
		return nil, err
	}

	return &returnViaCep, nil
}
