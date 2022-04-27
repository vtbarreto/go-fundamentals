package main

import "fmt"

type carro struct {
	nome            string
	valocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.valocidadeAtual = 0
	f.turboLigado = true

	fmt.Println(f.nome, f.valocidadeAtual)
	fmt.Println(f)
}
