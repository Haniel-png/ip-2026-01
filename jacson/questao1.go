package main

import "fmt"

func main() {

	var (
		a, b, c float64
	)

	fmt.Println("Escreva os três lados")
	fmt.Scan(&a, &b, &c)

	if (a+b > c) && (a+c > b) && (b+c > a) {
		fmt.Println("É um triângulo válido")
	} else {
		fmt.Println("É um triângulo inválido")
	}

	if (a != b) && (a != c) && (b != c) {
		fmt.Println("É um triÂngulo Escaleno")
	} else if (a == b) || (a == c) || (b == c) {
		fmt.Println("É um triângulo isósceles")
	} else if (a == b) && (b == c) {
		fmt.Println("É um triângulo equilátero")
	}
}