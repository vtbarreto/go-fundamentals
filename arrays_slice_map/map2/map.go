package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose":    11325,
		"Gabriel": 1234,
		"Pedro":   1200.1,
	}

	funcsESalarios["Rafaela Filho"] = 12350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
