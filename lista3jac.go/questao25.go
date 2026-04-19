package main

import (
	"fmt"
)

func main() {
	var soma float64 = 0.0
	
	var numerador float64 = 1.0
	var sinal float64 = 1.0

	fmt.Println("--- Calculando a Série S ---")

	for baseDenominador := 15.0; baseDenominador >= 1.0; baseDenominador-- {
		
		denominador := baseDenominador * baseDenominador
		
		termo := (numerador / denominador) * sinal
		
		soma += termo

		
		if sinal > 0 {
			fmt.Printf("+ (%.0f / %.0f) = %.4f\n", numerador, denominador, termo)
		} else {
			fmt.Printf("- (%.0f / %.0f) = %.4f\n", numerador, denominador, termo)
		}

		numerador *= 2.0      
		sinal *= -1.0         
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("O resultado final de S é: %.4f\n", soma)
}