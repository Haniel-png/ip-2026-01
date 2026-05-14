package main

import "fmt"

func maior(a, b, c int) int {
	maior := a
	if b > maior {
		maior = b
	}	
	if c > maior {
		maior = c
	}
	return maior
}
func main() {
	var a, b, c int
	fmt.Print("Digite 3 números: ")
	fmt.Scan(&a, &b, &c)
	resultado := maior(a, b, c)
	fmt.Printf("O maior número é: %d\n", resultado)
}