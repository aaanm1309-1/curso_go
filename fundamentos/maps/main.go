package main

import "fmt"

func main() {

	println("------------------------------------------------")
	salarios := map[string]float32{"Adriano": 2999.59, "John": 1567.78, "Mart": 1000.5}

	fmt.Println(salarios["Adriano"])
	// fmt.Printf("len=%d cap%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)

	for i, v := range salarios {
		// println("Array:", i, " valor: ", v)
		// fmt.Printf("Array: %v e o valor é %v\n", i, v)
		fmt.Printf("Array: %v e o valor é %v\n", i, v)
	}

	delete(salarios, "John")

	salarios["Jo"] = 2388.97
	println("------------------------------------------------")

	for i, v := range salarios {
		// println("Array:", i, " valor: ", v)
		// fmt.Printf("Array: %v e o valor é %v\n", i, v)
		fmt.Printf("Array: %v e o valor é %v\n", i, v)
	}

	println("------------------------------------------------")

	sal := make(map[string]int)
	sal1 := map[string]int{}

	sal["Olha"] = 2584

	println(sal["Olha"])
	println("------------------------------------------------")

	sal1["Olha1"] = 24444
	sal1["Olha2"] = 2333

	println(sal1["Olha2"])
	println("------------------------------------------------")

	for _, v := range salarios {
		// println("Array:", i, " valor: ", v)
		// fmt.Printf("Array: %v e o valor é %v\n", i, v)
		fmt.Printf("O valor do salário é %v\n", v)
	}

	println("------------------------------------------------")

	// fmt.Printf("len=%d cap%d %v\n", len(meuSlice[:0]), cap(meuSlice[:0]), meuSlice[:0])

	// fmt.Printf("len=%d cap%d %v\n", len(meuSlice[:4]), cap(meuSlice[:4]), meuSlice[:4])

	// fmt.Printf("len=%d cap%d %v\n", len(meuSlice[2:]), cap(meuSlice[2:]), meuSlice[2:])

	// fmt.Printf("len=%d cap%d %v\n", len(meuSlice[4:6]), cap(meuSlice[4:6]), meuSlice[4:6])

	// fmt.Printf("len=%d cap%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)

	// meuSlice = append(meuSlice, 201, 155)

	// fmt.Printf("len=%d cap%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)

}
