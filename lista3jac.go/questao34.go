package main

import (
	"fmt"
)

func main() {
	var n1, n2 int

	fmt.Println("--- Calculadora de MMC ---")
	fmt.Print("Digite o primeiro número (N1): ")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número (N2): ")
	fmt.Scan(&n2)

	if n1 <= 0 || n2 <= 0 {
		fmt.Println("Erro: Por favor, digite apenas números inteiros positivos.")
		return
	}

	maior := n1
	if n2 > n1 {
		maior = n2
	}

	mmc := maior

	for {

		if mmc%n1 == 0 && mmc%n2 == 0 {
			break 
		}
		
		mmc += maior 
	}

	fmt.Printf("\nO Mínimo Múltiplo Comum (MMC) entre %d e %d é: %d\n", n1, n2, mmc)
}