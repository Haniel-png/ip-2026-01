package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Printf("Digite um número inteiro entre 5 e 2000: ")
	fmt.Scanln(&N)
	if N >= 5 && N <= 2000 {
		for i := 2; i <= N; i += 2 {
			fmt.Printf("%d^2 = %d\n", i, i*i)
		}
	} else if N < 5 || N > 2000 {
		fmt.Printf("Número inválido. Digite um número inteiro entre 5 e 2000.")
	}

}
