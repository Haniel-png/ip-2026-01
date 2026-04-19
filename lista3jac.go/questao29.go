package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Digite um numero positvo e inteiro:")
	fmt.Scan(&N)
	if N < 0 {
		fmt.Print("Erro, digite um numero positvo e inteiro:")
		return
	}
	soma := 0
	for i := 1; i <= N; i++ {
		soma = soma + i
	}
	fmt.Printf("A soma dos números inteiros de 1 a N é %d", soma)
}
