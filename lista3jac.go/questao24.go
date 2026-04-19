package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Tabela de Aproximação do Seno (Série de Maclaurin) ---")
	fmt.Println(" Ângulo (A)  |   Sen(A) Aproximado")
	fmt.Println("----------------------------------------------------------")

	for i := 0; i <= 63; i++ {
		A := float64(i) * 0.1
		A3 := A * A * A
		A5 := A3 * A * A
		A7 := A5 * A * A

		senoA := A - (A3 / 6.0) + (A5 / 120.0) - (A7 / 5040.0)
		fmt.Printf(" %4.1f rad    |   %11.6f\n", A, senoA)
	}
	fmt.Println("---------------------------------------------------------------------------------------")
}