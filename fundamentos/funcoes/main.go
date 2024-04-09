package main

import (
	"errors"
	"fmt"
)

func main() {

	// fmt.Println("Retorno da Funcao: ", soma(1, 3))
	fmt.Println(soma(1, 3))
	println("------------------------------------------------")

	fmt.Println(soma(15, 39))
	println("------------------------------------------------")

	valor, err := somaComError(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

	println("------------------------------------------------")

	valor1, err1 := somaComError(15, 39)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(valor1)
	println("------------------------------------------------")

}

func soma(a, b int) (int, bool) {
	if (a + b) >= 50 {
		return a + b, true
	}
	return a + b, false
}

func somaComError(a, b int) (int, error) {
	if (a + b) < 50 {
		return -1, errors.New("Valor deve ser maior ou igual a 50")
	}
	return a + b, nil
}
