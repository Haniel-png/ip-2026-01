package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64
	var Ab, v float64

	fmt.Print("Digite a altura e o lado do triângulo equilátero: ")
	fmt.Scan(&h, &a)


	Ab = (3 * a * a * math.Sqrt(3)) / 2

	
	v = (Ab * h) / 3

	
	fmt.Printf("O VOLUME E = %.2f\n", v)
}