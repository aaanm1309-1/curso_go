package main

import "fmt"

const mensagem = "Oi meu amigo adriano"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Adriano"
	e float64 = 1.5
	f ID      = 1
)

func main() {
	a := "b"
	fmt.Printf("O tipo de A é %T e o valor é %v", a, a)
	println()
	fmt.Printf("O tipo de B é %T e o valor é %v", b, b)
	println()
	fmt.Printf("O tipo de C é %T e o valor é %v", c, c)
	println()
	fmt.Printf("O tipo de D é %T e o valor é %v", d, d)
	println()
	fmt.Printf("O tipo de E é %T e o valor é %v", e, e)
	println()
	fmt.Printf("O tipo de F é %T e o valor é %v", f, f)
	println()
	x()

}

func x() {
	fmt.Println("Imprimir funcao X")
}
