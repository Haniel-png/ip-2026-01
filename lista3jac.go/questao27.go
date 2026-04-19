package main

import (
	"fmt"
	"math"
)

func fatorial(n int) float64 {
	if n == 0 {
		return 1.0
	}
	fat := 1.0
	for i := 1; i <= n; i++ {
		fat *= float64(i)
	}
	return fat
}

func main() {
	var x float64

	fmt.Println("--- Calculadora de Cosseno usando math.Pow ---")
	fmt.Print("Digite o valor do ângulo x (em radianos): ")
	fmt.Scan(&x)


	soma := 0.0
	sinal := 1.0 

	for i := 0; i < 20; i++ {
		
		expoente := float64(2 * i)

		numerador := math.Pow(x, expoente)
		
		denominador := fatorial(int(expoente))

		termo := (numerador / denominador) * sinal
		
		soma += termo

		sinal *= -1.0
	}

	cosReal := math.Cos(x)

	diferenca := math.Abs(cosReal - soma)

	
	fmt.Println("\n--- Resultados ---")
	fmt.Printf("a) Cosseno calculado (20 termos com math.Pow): %.15f\n", soma)
	fmt.Printf("b) Cosseno fornecido pela função math.Cos(X):  %.15f\n", cosReal)
	fmt.Printf("c) Diferença (Erro de aproximação):            %.15f\n", diferenca)
}