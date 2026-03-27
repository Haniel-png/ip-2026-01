package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C, delta float64


	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&A)	
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&B)	
	fmt.Print("Digite o valor de C: ")
	fmt.Scan(&C)

	delta = B*B - 4*A*C
	if delta < 0 {
		fmt.Print("A equação possui raízes imaginárias.")
	} else if delta == 0 {
		x := -B / (2 * A)
		fmt.Printf("A equação possui uma raiz real: %.2f", x)
	} else {
		x1 := (-B + math.Sqrt(delta)) / (2 * A)
		x2 := (-B - math.Sqrt(delta)) / (2 * A)
		fmt.Printf("A equação possui duas raízes reais: %.2f e %.2f", x1, x2)
	}
}