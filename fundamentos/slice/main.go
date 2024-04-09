package main

import "fmt"

func main() {

	meuSlice := []int{4, 5, 6, 7, 54, 32, 87, 99}

	fmt.Println(len(meuSlice))
	fmt.Printf("len=%d cap%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)

	for i, v := range meuSlice {
		// println("Array:", i, " valor: ", v)
		// fmt.Printf("Array: %v e o valor é %v\n", i, v)
		fmt.Printf("Array: %d e o valor é %d\n", i, v)
	}

	fmt.Printf("len=%d cap%d %v\n", len(meuSlice[:0]), cap(meuSlice[:0]), meuSlice[:0])

	fmt.Printf("len=%d cap%d %v\n", len(meuSlice[:4]), cap(meuSlice[:4]), meuSlice[:4])

	fmt.Printf("len=%d cap%d %v\n", len(meuSlice[2:]), cap(meuSlice[2:]), meuSlice[2:])

	fmt.Printf("len=%d cap%d %v\n", len(meuSlice[4:6]), cap(meuSlice[4:6]), meuSlice[4:6])

	fmt.Printf("len=%d cap%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)

	meuSlice = append(meuSlice, 201, 155)

	fmt.Printf("len=%d cap%d %v\n", len(meuSlice), cap(meuSlice), meuSlice)

}
