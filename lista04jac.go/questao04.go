package main

import (
	"fmt"
	"strconv"
)

func decimalParaBinario(n int) string {
	
	if n == 0 {
		return "0"
	}
	
	if n == 1 {
		return "1"
	}

	resto := strconv.Itoa(n % 2)
	

	return decimalParaBinario(n/2) + resto
}

func main() {
	var decimal int

	fmt.Println("--- Conversor Decimal para Binário ---")
	fmt.Print("Digite um número inteiro decimal (positivo): ")
	fmt.Scan(&decimal)

	if decimal < 0 {
		fmt.Println("Erro: Este algoritmo simples processa apenas números positivos.")
		return
	}

	binario := decimalParaBinario(decimal)

	fmt.Println("--------------------------------------")
	fmt.Printf("O número decimal %d em binário é: %s\n", decimal, binario)
}