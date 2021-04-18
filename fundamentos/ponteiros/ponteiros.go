package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // Pegando o endereco da variavel

	*p++
	i++

	//Go nao tem aritmetica de ponteiros
	//p++

	fmt.Println(p, *p, i, &i)
}
