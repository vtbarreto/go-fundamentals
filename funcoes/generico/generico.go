package main

import "fmt"

func Sum[V int64 | float64](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := []int64{1, 2, 3, 4, 5, 6}
	floats := []float64{12.3, 13.5, 10.7, 11.42}
	fmt.Println("soma int 64:", Sum[int64](ints))
	fmt.Println("soma float 64:", Sum[float64](floats))
}
