package main

import "fmt"

func main() {
	var soma, qtd int
	var media float64

	soma = 0
	qtd = 0

	for i := 50 ; i <= 70; i = i + 2 {
		soma = soma + i
	}
	fmt.Printf("A soma dos números pares entre 50 e 70 é: %d\n", soma)

	for i := 50 ; i <= 70 ; i = i + 2 {
		qtd = qtd + 1
	}
	media = float64(soma) / float64(qtd)
	fmt.Printf("A média dos números pares entre 50 e 70 é: %.2f\n", media)
}