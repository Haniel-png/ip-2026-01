package main

import (
	"fmt"
)

func inverteArray(numeros []int, n int) {
	if n <= 1 {
		return
	}

	numeros[0], numeros[n-1] = numeros[n-1], numeros[0]

	inverteArray(numeros[1:n-1], n-2)
}

func main() {
	valores := []int{10, 20, 30, 40, 50, 60}
	tamanho := len(valores)

	fmt.Println("--- Inversão de Array com Recursão ---")
	fmt.Printf("Array Original: %v\n", valores)


	inverteArray(valores, tamanho)

	fmt.Println("--------------------------------------")
	fmt.Printf("Array Invertido: %v\n", valores)
}