package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//mapas devems er inicializados
	aprovados := make(map[int]string)

	aprovados[223213] = "Maria"
	aprovados[123213] = "Pedro"
	aprovados[123233] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println(cpf, nome)
	}

	fmt.Println(aprovados[223213])
	delete(aprovados, 223213)
	fmt.Println(aprovados[223213])
}
