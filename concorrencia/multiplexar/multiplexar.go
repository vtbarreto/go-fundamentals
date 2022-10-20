package main

import (
	"fmt"

	"github.com/viniciusteixeirabarreto/gohtml"
)

func encaminhar(origem <-chan string, destino chan string) {
	//infinitamente ira jogar de um canal para o outro

	//mas o for fica esperando o origem receber valor para rodar, nao fica rodando a todo momento
	for {
		destino <- <-origem
	}
}

//misturar mensagens em um canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string, 4)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		gohtml.Titulo("https://www.cod3r.com.br", "https://www.google.com.br"),
		gohtml.Titulo("https://www.bing.com", "https://www.youtube.com.br"),
	)

	//For vai rodando sozinho conforme as mensagens forem atribuidas ao canal
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}

}
