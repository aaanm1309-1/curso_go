package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero    int    `json:"numero,omitempty"`
	Saldo     int    `json:"saldo,omitempty"`
	Descricao string `json:"-"`
}

func main() {

	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"Numero": 2, "Saldo": 3300}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println("Numero da conta: ", contaX.Numero, " com saldo: ", contaX.Saldo)

	jsonPuroSemRelacao := []byte(`{"n": 2, "s": 3300}`)
	var contaX2 Conta
	err = json.Unmarshal(jsonPuroSemRelacao, &contaX2)
	if err != nil {
		panic(err)
	}
	println("Numero da conta 2 : ", contaX2.Numero, " com saldo: ", contaX2.Saldo)

}
