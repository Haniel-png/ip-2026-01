package main

import (
	"fmt"
)

func main() {
	var montanteCarlos float64

	fmt.Print("Digite o salário do Carlos: R$ ")
	fmt.Scan(&montanteCarlos)

	montanteJoao := montanteCarlos / 3.0

	meses := 0

	for montanteJoao < montanteCarlos {
		
		montanteCarlos += montanteCarlos * 0.02
		
		montanteJoao += montanteJoao * 0.05
		
		meses++
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Foram necessários %d meses para João alcançar ou ultrapassar Carlos.\n", meses)
	fmt.Printf("Valor final do Carlos: R$ %.2f\n", montanteCarlos)
	fmt.Printf("Valor final do João: R$ %.2f\n", montanteJoao)
}