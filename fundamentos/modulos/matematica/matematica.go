package matematica

func Soma[T int | float64](listaNumeros ...T) T {
	var valorResult T
	for _, valor := range listaNumeros {
		valorResult += valor
	}
	return valorResult
}
