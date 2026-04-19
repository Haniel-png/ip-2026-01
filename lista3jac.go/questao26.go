package main

import (
	"fmt"
)

func main() {
	var soma, termo float64 = 0.0, 0.0
	numerador := 100.0
	fatorial := 1.0
	for i := 0; i < 20; i++ {
		if i > 0 {
			fatorial = fatorial * float64(i)
		}
		numerador = 100 - float64(i)
		termo = numerador / fatorial
		soma = soma + termo
	}
	fmt.Printf("%f", soma)
}



