package main

import (
	"fmt"
	"time"
)

//Chanel é a forma de comucacao entre goroutines
//channel é um tipoe

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c //Recebendo dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)

}
