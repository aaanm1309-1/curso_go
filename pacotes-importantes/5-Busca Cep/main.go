package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, url := range os.Args[1:] {
		// println(url)
		req, err := http.Get("http://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisicão: %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
		}

		// fmt.Println(data)

		file, err := os.Create("resultadoCep.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf(
			`{"CEP": "%s", "Logradouro": "%s","Complemento": "%s","Bairro": "%s", "Localidade": "%s", "UF": "%s", "DDD": %s}`,
			data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Ddd))

		fmt.Printf(
			`{"CEP": "%s", "Logradouro": "%s", "Complemento": "%s", "Bairro": "%s", "Localidade": "%s", "UF": "%s", "DDD": %s}`,
			data.Cep, data.Logradouro, data.Complemento, data.Bairro, data.Localidade, data.Uf, data.Ddd)
	}

}
