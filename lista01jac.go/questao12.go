package main

import (
	"fmt"
)

func main() {
	var h int

	fmt.Scanln(&h)

	blocos := h / 3
	resto := h % 3

	valor := float64(blocos*10 + resto*5)

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}