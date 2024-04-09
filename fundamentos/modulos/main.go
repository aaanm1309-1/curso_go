package main

import (
	"fmt"

	"github.com/aaanm1309-1/curso_go/matematica"
	"github.com/google/uuid"
)

func main() {

	soma := matematica.Soma(10.5, 20.7, 30)
	fmt.Println("Resultado: ", soma)
	fmt.Println("UUID: ", uuid.New())
}
