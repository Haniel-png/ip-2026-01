package main

import (
	"fmt"
)

func main() {
	var tipo string
	var consumo, valor float64
	
	fmt.Print("Digite o tipo de consumo (R, C ou I): ")
	fmt.Scan(&tipo)
	fmt.Print("Digite o consumo em m³: ")
	fmt.Scan(&consumo)

	switch tipo {
		
	case "R", "r":	
	 valor = 5 + 0.05 * consumo
		
	case "C", "c":
	if consumo <= 80 {
		valor = 500
	} else { 
		valor = 500 + ( 0.25 * (consumo - 80))
	}	
	
    case "I", "i":
	if consumo <= 100 {
		valor = 800	
	} else {
		valor = 800 + ( 0.04 * (consumo - 100))
	}	
	default:
		fmt.Print("Tipo de consumo inválido.")
		return
	}		
	fmt.Printf("O valor a ser pago é: R$ %.2f", valor)			
}