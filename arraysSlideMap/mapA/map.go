package main

import "fmt"

func main() {
	porLetra := map[string]map[string]float64{
		"G": {
			"Gabiela Silva": 15456,
			"Guga":          12321,
		},
		"J": {
			"Jose": 123,
		},
		"P": {
			"Pedrinho": 123,
		},
	}

	delete(porLetra, "P")

	for letra, funcs := range porLetra {
		fmt.Println(letra, funcs)
	}
}
