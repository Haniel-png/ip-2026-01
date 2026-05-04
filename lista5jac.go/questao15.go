package main

import "fmt"

func main() {
	var v1 [30]int
	var v2 [30]int

	
	fmt.Println("Digite 30 números inteiros:")
	for i := 0; i < 30; i++ {
		fmt.Scan(&v1[i])
	}

	
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
		
			v2[i] = v1[i] * 2
		} else {
			v2[i] = v1[i] * 3
		}
	}

	
	fmt.Println("\nVetor resultante:")
	for i := 0; i < 30; i++ {
		fmt.Printf("%d ", v2[i])
	}
	fmt.Println()
}