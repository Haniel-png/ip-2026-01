package main

import (
	"fmt"
)

func main() {
	var octal int

	fmt.Println("--- Conversor de Base 8 para Base 10 ---")
	fmt.Print("Digite um número inteiro positivo na base 8: ")
	fmt.Scan(&octal)

	if octal < 0 {
		fmt.Println("Erro: O programa aceita apenas números positivos.")
		return
	}

	numeroAtual := octal
	decimal := 0
	multiplicador := 1 

	for numeroAtual > 0 {
		digito := numeroAtual % 10 

		if digito >= 8 {
			fmt.Printf("Erro: O número digitado não é um octal válido (contém o dígito %d).\n", digito)
			return
		}

			decimal += digito * multiplicador

			multiplicador *= 8             
		numeroAtual = numeroAtual / 10 
	}

	fmt.Printf("\nO equivalente de %d (base 8) na base 10 é: %d\n", octal, decimal)
}