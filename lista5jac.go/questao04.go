package main

import (
	"fmt"
)

func main() {
	var A [10]int
	
	
	frequencias := make(map[int]int)
	
	temRepetido := false

	fmt.Println("--- Leitura do Vetor A ---")
	
	
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&A[i])

		frequencias[A[i]]++
	}

	fmt.Println("\n--- Elementos Repetidos ---")
	
	for numero, quantidade := range frequencias {
		
		if quantidade > 1 {
			fmt.Printf("O número %d se repete %d vezes.\n", numero, quantidade)
			temRepetido = true // Sinaliza que achamos pelo menos um repetido
		}
	}

	if !temRepetido {
		fmt.Println("Não há nenhum elemento repetido no vetor.")
	}
}