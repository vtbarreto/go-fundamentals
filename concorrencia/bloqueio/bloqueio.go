package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operacao bloqueante
	fmt.Println("So depois que for lido")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c) //operacao bloqueante //Ele esperou a funcao ROTINA atribuir valor ao channel ou seja vai demorar um seguro
	fmt.Println("Foi lido")
	fmt.Println(<-c) //DeadLock //Como o codigo sabe que nao vai ter mais dado a ser enviado, ele da erro no codigo de bloqueio eterno
}
