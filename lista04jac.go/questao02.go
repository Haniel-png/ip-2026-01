package main

import (
	"fmt"
)

func somaArray(numeros []float64) float64 {
	if len(numeros) == 0 {
		return 0.0
	}

	return numeros[0] + somaArray(numeros[1:])
}

func main() {
	valores := []float64{2.5, 3.5, 6.0, 8.0}

	fmt.Println("--- Somatório Recursivo de Array ---")
	fmt.Printf("Valores no array: %v\n", valores)


	total := somaArray(valores)

	fmt.Println("------------------------------------")
	fmt.Printf("A soma de todos os valores é: %g\n", total)
}