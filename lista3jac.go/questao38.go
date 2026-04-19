package main

import "fmt"

func main() {
	var cpf [11]int

	fmt.Println("Digite os 11 dígitos do CPF separados por espaço:")

	for i := 0; i < 11; i++ {
		fmt.Scan(&cpf[i])
	}

	soma := 0
	peso := 10

	for i := 0; i < 9; i++ {
		soma += cpf[i] * peso
		peso--
	}

	resto := soma % 11
	var dig1 int

	if resto < 2 {
		dig1 = 0
	} else {
		dig1 = 11 - resto
	}


	soma = 0
	peso = 11

	for i := 0; i < 10; i++ {
		soma += cpf[i] * peso
		peso--
	}

	resto = soma % 11
	var dig2 int

	if resto < 2 {
		dig2 = 0
	} else {
		dig2 = 11 - resto
	}

	if dig1 == cpf[9] && dig2 == cpf[10] {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}