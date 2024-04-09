package main

import "fmt"

type PessoaGeneric interface{}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
	Ativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	EnderecoCliente Endereco
}

type Empresa struct {
	Nome            string
	TempoDeAbertura int
	Ativo           bool
	Endereco
	EnderecoCliente Endereco
}

func (c PessoaGeneric) andou() {

}

func (c *Cliente) Desativar() {
	// func (c *Cliente) Desativar() {

	c.Ativo = false
	// return c
}

func (c *Cliente) Ativar() {
	c.Ativo = true
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func Ativacao(pessoa Pessoa) {
	pessoa.Ativar()
}

func main() {

	novoCliente := Cliente{
		Nome:  "Adriano",
		Idade: 49,
		Ativo: true,
	}

	novoCliente.Ativo = false
	novoCliente.EnderecoCliente.Logradouro = "Enredeco do Cliente"
	novoCliente.Logradouro = "Enredeco do Geral Cliente"
	novoCliente.Numero = 10
	novoCliente.Cidade = "Cidade Geral do Cliente"
	novoCliente.Estado = "Estado Geral do Cliente"

	fmt.Printf("O nome do cliente é %s.\n", novoCliente.Nome)
	fmt.Println("--------------------------------")
	fmt.Printf("Nome %s, Idade: %d, Ativo: %t.\n",
		novoCliente.Nome,
		novoCliente.Idade,
		novoCliente.Ativo,
	)
	fmt.Printf("Endereco %s, Endereco Cliente: %s.\n",
		novoCliente.Logradouro,
		novoCliente.EnderecoCliente.Logradouro,
	)
	fmt.Println("--------------------------------")
	fmt.Println("Endereco ",
		novoCliente.Endereco,
	)
	fmt.Println("--------------------------------")

	novoCliente.Ativar()

	fmt.Println("--------Ativar------------------")
	fmt.Printf("Nome %s, Idade: %d, Ativo: %t.\n",
		novoCliente.Nome,
		novoCliente.Idade,
		novoCliente.Ativo,
	)
	fmt.Println("--------------------------------")

	// novoCliente.Ativo = true
	novoCliente.Desativar()

	fmt.Println("-------Desativar----------------")
	fmt.Printf("Nome %s, Idade: %d, Ativo: %t.\n",
		novoCliente.Nome,
		novoCliente.Idade,
		novoCliente.Ativo,
	)
	fmt.Println("--------------------------------")

	Ativacao(&novoCliente)
	fmt.Println("--------Ativacao---------------")
	fmt.Printf("Nome %s, Idade: %d, Ativo: %t.\n",
		novoCliente.Nome,
		novoCliente.Idade,
		novoCliente.Ativo,
	)
	fmt.Println("--------------------------------")

	Desativacao(&novoCliente)
	fmt.Println("--------Desativacao-------------")
	fmt.Printf("Nome %s, Idade: %d, Ativo: %t.\n",
		novoCliente.Nome,
		novoCliente.Idade,
		novoCliente.Ativo,
	)
	fmt.Println("--------------------------------")

}
