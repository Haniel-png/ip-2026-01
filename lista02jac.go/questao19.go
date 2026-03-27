package main

import (
	"fmt"
	"math"
)

func main() {
	var tipo string
	var r, h, area, volume float64
	const pi = 3.14

	fmt.Print("Digite o tipo de figura (1 para cone, 2 para cilindro ou 3 para esfera): ")
	fmt.Scan(&tipo)

	switch tipo {
	case "1":
		fmt.Print("Digite o raio do cone: ")
		fmt.Scan(&r)
		fmt.Print("Digite a altura do cone: ")
		fmt.Scan(&h)
		area = pi * r * (r + math.Sqrt(h*h+r*r))
		volume = (1.0 / 3.0) * pi * r * r * h

	case "2":
		fmt.Print("Digite o raio do cilindro: ")
		fmt.Scan(&r)
		fmt.Print("Digite a altura do cilindro: ")
		fmt.Scan(&h)
		area = 2 * pi * r * (r + h)
		volume = pi * r * r * h
	case "3":
		fmt.Print("Digite o raio da esfera: ")
		fmt.Scan(&r)
		area = 4 * pi * r * r
		volume = (4.0 / 3.0) * pi * r * r * r
	default:
		fmt.Println("Tipo de figura inválido!")
		return
	}
	fmt.Printf("A área da figura é: %.2f", area)
	fmt.Printf("O volume da figura é: %.2f", volume)
}