package main

import (
	"fmt"
)

func main() {
	var id int
	var n1, n2, n3, mediaEx, mediaFinal float64
	var conceito string

	fmt.Print("Digite o número do aluno: ")
	fmt.Scan(&id)

	fmt.Print("Digite a nota 1: ")
	fmt.Scan(&n1)

	fmt.Print("Digite a nota 2: ")
	fmt.Scan(&n2)

	fmt.Print("Digite a nota 3: ")
	fmt.Scan(&n3)

	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scan(&mediaEx)

	mediaFinal = (n1 + n2*2 + n3*3 + mediaEx) / 7

	switch {
	case mediaFinal >= 9 && mediaFinal <= 10:
		conceito = "A"
	case mediaFinal >= 7.5:
		conceito = "B"
	case mediaFinal >= 6:
		conceito = "C"
	case mediaFinal >= 4:
		conceito = "D"
	default:
		conceito = "E"
	}

	var situacao string
	if conceito == "A" || conceito == "B" || conceito == "C" {
		situacao = "APROVADO"
	} else {
		situacao = "REPROVADO"
	}

	fmt.Println("\n--- Resultado ---")
	fmt.Printf("Aluno: %d\n", id)
	fmt.Printf("Notas: %.2f, %.2f, %.2f\n", n1, n2, n3)
	fmt.Printf("Média dos exercícios: %.2f\n", mediaEx)
	fmt.Printf("Média final: %.2f\n", mediaFinal)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Situação: %s\n", situacao)
}