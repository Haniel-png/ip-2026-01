package main

import (
	"fmt"
)

func main() {
	var n float64

	fmt.Printf("Digite um número (entre 0 e 10): ")
	fmt.Scan(&n)
	if n >= 0 && n < 6 {
		fmt.Printf("Conceito: D")
	} else if n >= 6 && n < 7.5 {
		fmt.Printf("Conceito: C")
	} else if n >= 7.5 && n < 9 {
		fmt.Printf("Conceito: B")
	} else if n >= 9 && n <= 10 {
		fmt.Printf("Conceito: A")
	} else {
		fmt.Printf("Número inválido. Digite um número entre 0 e 10.")
	}
}