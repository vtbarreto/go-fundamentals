package main

import (
	"fmt"
	"time"

	"github.com/viniciusteixeirabarreto/gohtml"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := gohtml.Titulo(url1)
	c2 := gohtml.Titulo(url2)
	c3 := gohtml.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Deu ruim"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.bing.com",
		"https://www.google.com",
		"https://www.udemy.com",
	)

	fmt.Println(campeao)
}
