package main

import (
	"fmt"
)

func main() {
	var r, h float64
	const pi = 3.14159
	const custoPorMetro = 100.0

	
	fmt.Scanln(&r)
	fmt.Scanln(&h)

	
	ac := pi * r * r
	al := 2 * pi * r * h
	at := 2*ac + al
	custo := at * custoPorMetro


	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}