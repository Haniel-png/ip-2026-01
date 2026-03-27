package main

import (
	"fmt"
)

func main() {
	var preco float64
	var codigo int
	var valorFinal float64

	fmt.Print("Digite o preço do produto: R$ ")
	fmt.Scan(&preco)

	fmt.Println("\nEscolha a condição de pagamento:")
	fmt.Println("1 - À vista (dinheiro/cheque) - 10% de desconto")
	fmt.Println("2 - À vista (cartão de crédito) - 5% de desconto")
	fmt.Println("3 - Em 2x - sem juros")
	fmt.Println("4 - Em 3x - 10% de juros")
	fmt.Print("Opção: ")
	fmt.Scan(&codigo)

	switch codigo {
	case 1:
		valorFinal = preco * 0.90
	case 2:
		valorFinal = preco * 0.95
	case 3:
		valorFinal = preco
	case 4:
		valorFinal = preco * 1.10
	default:
		fmt.Println("Código inválido!")
		return
	}

	fmt.Printf("\nValor final: R$ %.2f\n", valorFinal)

	if codigo == 3 {
		fmt.Printf("Parcelado em 2x de R$ %.2f\n", valorFinal/2)
	} else if codigo == 4 {
		fmt.Printf("Parcelado em 3x de R$ %.2f\n", valorFinal/3)
	}
}