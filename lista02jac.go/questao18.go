package main

import (
	"fmt"
)

func main() {
	var preco, precoFinal float64
	var categoria string
	var dia int
	

	fmt.Print("Digite o preço normal do DVD: R$ ")
	fmt.Scan(&preco)

	fmt.Print("Digite a categoria (C - comum, L - lançamento): ")
	fmt.Scan(&categoria)

	fmt.Print("Digite o dia da semana (1=Dom, 2=Seg, ..., 7=Sáb): ")
	fmt.Scan(&dia)

	precoFinal = preco

	if dia == 2 || dia == 3 || dia == 5 {
		precoFinal *= 0.60 
	}

	if categoria == "L" || categoria == "l" {
		precoFinal *= 1.15 
	}

	if dia < 1 || dia > 7 {
		fmt.Println("Dia da semana inválido!")
		return
	}

	fmt.Printf("\nPreço final do aluguel: R$ %.2f\n", precoFinal)
}