package main

import (
	"fmt"
)

func main() {
	var x, y int
	
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&x)
	y = 0
	for y*y < x {
		y++
	}
	if y*y == x {
		fmt.Printf("O número %d é um quadrado perfeito.\n", x)
	} else { fmt.Printf("O número %d não é um quadrado perfeito.\n", x) }
}