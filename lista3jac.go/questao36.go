package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro e positivo:")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Print("Digite um numero positivo")
		return
	}
	if n == 0 {
		fmt.Printf("O valor na base binária é 0")
		return
	}
	hexMapas := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	var restos []string
	numeroAtual := n
	for numeroAtual > 0 {
		resto := numeroAtual % 16
		restos = append(restos, hexMapas[resto])
		numeroAtual = numeroAtual / 16
	}
	for i := len(restos) - 1; i >= 0; i-- {
		fmt.Print(restos[i])
	}
	fmt.Println()
}
