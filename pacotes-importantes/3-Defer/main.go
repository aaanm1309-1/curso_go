package main

func main() {
	defer println("Primeira Linha...")
	defer println("Segunda Linha...")
	println("Terceira Linha...")
	println("Quarta Linha...")
	defer println("Quinta Linha...")
	println("Sexta Linha...")

}
