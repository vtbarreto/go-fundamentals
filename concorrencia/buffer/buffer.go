package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("Executou! 1")
	ch <- 2
	fmt.Println("Executou! 2")
	ch <- 3
	fmt.Println("Executou! 3")
	ch <- 4
	fmt.Println("Executou! 4")
	ch <- 5
	fmt.Println("Executou! 5")
	ch <- 6
	fmt.Println("Executou! 6")
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second) //Colocando pausa para nao processar enquanto o buffer eh carregado

	fmt.Println(<-ch)
}
