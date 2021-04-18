package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	compraTv50 := trab1 && trab2
	compraTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return compraTv50, compraTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}
