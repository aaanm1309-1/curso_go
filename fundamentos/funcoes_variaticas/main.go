package main

import (
	"fmt"
)

func main() {

	fmt.Println(soma(1, 3, 56, 70, 89, 54, 645654, 6456))

}

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
