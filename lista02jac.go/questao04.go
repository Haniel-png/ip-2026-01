package main

import (
	f "fmt"
	"math"
)

func main() {

	var numero, resultado float64

	f.Print("Digite um número:")
	f.Scan(&numero)

	if numero >= 0 {
		resultado = math.Sqrt(float64(numero))
		f.Printf("O resultado é: %.2f", resultado)
	} else {
		resultado = math.Pow(float64(numero), 2)
		f.Printf("O resultado é: %.2f", resultado)
	}

}
