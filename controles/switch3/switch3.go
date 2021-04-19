package main

import "fmt"

func tipo(i interface{}) string {

	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "nao sei"
	}

}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(2))
	fmt.Println(tipo("ola"))
}
