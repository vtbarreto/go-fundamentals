package main

import (
	"fmt"

	"github.com/Fundamentos-em-Go/mooncake/mocks"
	"github.com/GuilhermeCaruso/mooncake"
)

type Imprimivel interface {
	ToString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces sao implementadas implicitamente
func (p pessoa) ToString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) ToString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x Imprimivel) {
	fmt.Println(x.ToString())
}

func main() {
	var coisa Imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.ToString())
	imprimir(coisa)

	coisa = produto{"Calca", 79.90}
	fmt.Println(coisa.ToString())
	imprimir(coisa)

	p2 := produto{"Calcas Jeans", 179}
	imprimir(p2)

	a := mooncake.NewAgent()

	c := mocks.NewMockImprimivel(a)

	c.Prepare().ToString().SetReturn("Caruso").AnyTime()

	imprimir(c)
}
