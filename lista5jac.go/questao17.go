package main

import "fmt"


func ehPrimo(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var v [10]int

	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println("\nNúmeros primos e suas posições:")
	for i := 0; i < 10; i++ {
		if ehPrimo(v[i]) {
			fmt.Printf("Valor %d na posição %d\n", v[i], i)
		}
	}
}