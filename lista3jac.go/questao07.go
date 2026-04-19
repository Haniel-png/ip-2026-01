package main

import "fmt"

func main() {
	var num, soma, maior, menor, somapar, qtdpar, somaimpar, qtdimpar int
	qtd := 0
	soma = 0

	for {
		fmt.Print("Digite um número inteiro positivo): ")
		fmt.Scan(&num)
		if num == 30000 {
			break
		}
		qtd++
		soma = soma + num
		if qtd == 1 {
			maior = num
			menor = num
		} else {
			if num > maior {
				maior = num
			}
			if num < menor {
				menor = num
			}
			if num%2 == 0 {
				somapar = somapar + num
				qtdpar = qtdpar + 1
			} else {
				somaimpar = somaimpar + num
				qtdimpar = qtdimpar + 1
			}
		}

	}
	media := float64(soma) / float64(qtd)
	fmt.Printf("A média dos números digitados é: %.2f\n", media)
	fmt.Printf("O maior número digitado é: %d\n", maior)
	fmt.Printf("O menor número digitado é: %d\n", menor)
	if qtdpar > 0 {
		mediapar := float64(somapar) / float64(qtdpar)
		fmt.Printf("A média dos números pares digitados é: %.2f\n", mediapar)
	} else {
		fmt.Println("Nenhum número par foi digitado.")
	}
	if qtdimpar > 0 {
		mediaimpar := float64(somaimpar) / float64(qtdimpar)
		fmt.Printf("A média dos números ímpares digitados é: %.2f\n", mediaimpar)
	} else {
		fmt.Println("Nenhum número ímpar foi digitado.")
	}
}
