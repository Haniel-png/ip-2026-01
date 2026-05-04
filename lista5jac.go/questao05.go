package main

import (
	"fmt"
)

func main() {
	var vetor [10]int

	fmt.Println("--- Busca do Menor Elemento ---")

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
	}

	X := vetor[0] 
		P := 0        

	
	for i := 1; i < 10; i++ {
		if vetor[i] < X {
			X = vetor[i] 
			P = i        
		}
	}

	fmt.Println("\n-------------------------------")
	
	fmt.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", X, P)
}