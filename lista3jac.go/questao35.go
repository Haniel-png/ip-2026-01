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
	var restos []int
	numeroAtual := n
	for numeroAtual > 0 {
		resto := numeroAtual % 2
		restos = append(restos, resto)
		numeroAtual = numeroAtual/2
	} 
	for i:= len(restos) - 1; i>=0; i-- {
		fmt.Print(restos[i])
	}
	fmt.Println()
}