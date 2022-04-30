package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interacao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq vc nao fala comigo ?", 3)
	// fale("Joao", "So posso falar depois de vc ?", 3)

	// go fale("Maria", "Ei...", 4)
	// go fale("Joao", "Opa...", 4)

	go fale("Maria", "Salve...", 4)
	fale("Joao", "Opa...", 4)

}
