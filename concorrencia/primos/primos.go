package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 300)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c) //cap() capacidade de um channel
	for primo := range c {
		fmt.Printf("%d ", primo) //Esse codigo esta rodando em concorrencia conforme o channel for tendo valores atribuidos
	}
	fmt.Println("Fim!")
}
