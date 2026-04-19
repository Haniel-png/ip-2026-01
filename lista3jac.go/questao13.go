package main

import "fmt"

func main() {
	var soma float64 = 0.0
	var denominador float64 = 1.0
	var numerador float64 = 1.0
	var termo float64 = 0.0

	for denominador >= 1.0 && denominador <= 50.0 {
		fmt.Print("Calculando a soma")
		denominador += 1
		numerador += 2
		termo = (numerador / denominador)
		soma += termo
	}
	fmt.Print("A soma será igual a:", soma)
}
