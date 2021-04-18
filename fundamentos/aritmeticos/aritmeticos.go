package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtracao =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Modulo =>", a%b)

	//bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0

	//outras operacoes usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
}
