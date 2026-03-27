package main

import (
	"fmt"
)

func main() {
	var precoBase float64
	var opcao string
	var precoFinal float64

	fmt.Print("Digite o preço de fábrica do carro: R$ ")
	fmt.Scan(&precoBase)

	precoFinal = precoBase

	fmt.Print("Deseja ar condicionado? (s/n): ")
	fmt.Scan(&opcao)
	if opcao == "s" || opcao == "S" {
		precoFinal += 1750.00
	}

	fmt.Print("Deseja pintura metálica? (s/n): ")
	fmt.Scan(&opcao)
	if opcao == "s" || opcao == "S" {
		precoFinal += 800.00
	}

	fmt.Print("Deseja vidro elétrico? (s/n): ")
	fmt.Scan(&opcao)
	if opcao == "s" || opcao == "S" {
		precoFinal += 1200.00
	}

	fmt.Print("Deseja direção hidráulica? (s/n): ")
	fmt.Scan(&opcao)
	if opcao == "s" || opcao == "S" {
		precoFinal += 2000.00
	}

	fmt.Printf("\nPreço final do carro: R$ %.2f\n", precoFinal)
}