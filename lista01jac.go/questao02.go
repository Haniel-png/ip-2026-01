package main

import "fmt"

func main() {
	var ingressosVendidos, rendadojogo, porcentagem1, porcentagem2, porcentagem3, porcentagem4 float64

	fmt.Print("Digite o número de ingressos vendidos: ")
    fmt.Scanln(&ingressosVendidos)
	fmt.Print("Digite a porcentagem para a categoria Popular: ")
	fmt.Scanln(&porcentagem1)
	fmt.Print("Digite a porcentagem para a categoria Geral: ")
	fmt.Scanln(&porcentagem2)
	fmt.Print("Digite a porcentagem para a categoria Arquibancada: ")
	fmt.Scanln(&porcentagem3)
	fmt.Print("Digite a porcentagem para a categoria Cadeiras: ")
	fmt.Scanln(&porcentagem4)

	rendadojogo = (porcentagem1*1 + porcentagem2*5 + porcentagem3*10 + porcentagem4*20) * ingressosVendidos
	fmt.Printf("O valor arrecadado com os ingressos é: R$ %.2f", rendadojogo)
}
