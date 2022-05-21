package main

import (
	"fmt"

	gohtml "github.com/viniciusteixeirabarreto/gohtml"
)

func main() {
	t1 := gohtml.Titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := gohtml.Titulo("https://www.amazon.com.br", "https://www.youtube.com")

	fmt.Println(<-t1, "|")
	fmt.Println(<-t2, "|")

	fmt.Println(<-t1, "|")
	fmt.Println(<-t2, "|")
}
