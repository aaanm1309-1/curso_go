package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello BB")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso!: %d\n", tamanho)

	f2, err := os.Create("arquivoEmBytes.txt")
	if err != nil {
		panic(err)
	}

	tamanho2, err := f2.Write([]byte("Hello escrevendo em bytes"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso em bytes!: %d\n", tamanho2)

	f.Close()

	f2.Close()

	/// Leitura

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	arquivo2, err := os.Open("arquivoEmBytes.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)

	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	os.Remove("arquivoEmBytes.txt")
	os.Remove("arquivo.txt")

}
