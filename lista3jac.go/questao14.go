package main

import (
	"fmt"
)

func main() {
	var n1, n2 int

	fmt.Print("Digite o primeiro número (N1): ")
	fmt.Scan(&n1)

	fmt.Print("Digite o segundo número (N2): ")
	fmt.Scan(&n2)
	
	if n1 > n2 {
		n1, n2 = n2, n1
	}

	fmt.Printf("\nNúmeros primos entre %d e %d:\n", n1, n2)

	encontrouPrimo := false

	for num := n1; num <= n2; num++ {
		
		if num < 2 {
			continue
		}

		ehPrimo := true

		for i := 2; i <= num/2; i++ {
			
			if num%i == 0 {
				ehPrimo = false
				break 
			}
		}

		if ehPrimo {
			fmt.Printf("%d ", num)
			encontrouPrimo = true
		}
	}

	if !encontrouPrimo {
		fmt.Println("Nenhum número primo encontrado neste intervalo.")
	} else {
		fmt.Println() 
	}
}