package main

import (
	"fmt"
)

func main() {
	var n int
	var s, soma float64
	fmt.Printf("Digite um número inteiro: ")
	fmt.Scan(&n)
	soma = 0
	if n > 0 {
		for i := 1; i <= n; i++ {
			s = 1 / float64(i)
			soma = soma + s
		}
		fmt.Printf("Soma dos termos da série: %.2f\n", soma)
	} else {
		fmt.Printf("Número inválido. Digite um número inteiro positivo.")
	}
  		}
	
