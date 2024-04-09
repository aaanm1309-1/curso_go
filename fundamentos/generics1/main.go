package main

type myNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, valor := range m {
		soma += valor
	}
	return soma
}

// func Compara[T Number](a T, b T) bool {
func Compara[T comparable](a T, b T) bool {

	if a == b {
		return true
	}
	return false
}

func main() {

	mapInteiro := map[string]int{"A": 1000, "B": 2500, "C": 3900}
	mapFloat := map[string]float64{"A": 1000.20, "B": 2500.50, "C": 3900.10}

	mapInteiro2 := map[string]myNumber{"A": 1000, "B": 2500, "C": 3900}

	println(Soma(mapInteiro))
	println(Soma(mapFloat))
	println(Soma(mapInteiro2))

	println(Compara(10, 10.000000))

	// Desativacao(&novoCliente)
	// fmt.Println("--------Desativacao-------------")
	// fmt.Printf("Nome %s, Idade: %d, Ativo: %t.\n",
	// 	novoCliente.Nome,
	// 	novoCliente.Idade,
	// 	novoCliente.Ativo,
	// )
	// fmt.Println("--------------------------------")

}
