package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro eh", reflect.TypeOf(32000))

	//sem sinal uint8 uint16 uint32 uint 64
	var b byte = 255
	fmt.Println("O byte e", reflect.TypeOf(b))

	//com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int e", i1)
	fmt.Println("O tipo de i1 e", reflect.TypeOf(1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune e", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais
	var x float32 = 49.99
	fmt.Println("O x", reflect.TypeOf(x))
	fmt.Println("O 49.99", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O bo e", reflect.TypeOf(bo))
	fmt.Println(bo)
	fmt.Println(!bo)

	//string
	s1 := "Ola meu nome e leo"
	fmt.Println("O s1 e", reflect.TypeOf(bo))
	fmt.Println(s1)
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string", len(s1))

	//string com multiplas linhas
	s2 := `Ola
	meu 
	friend`
	fmt.Println("O tamanho da string e", len(s2))

	//char ???
	char := 'a'
	fmt.Println("Otiupo de char e", reflect.TypeOf(char))
	fmt.Println(char)

}
