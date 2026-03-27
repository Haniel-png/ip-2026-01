package main

import "fmt"

func main() {
	
	var n, a, b int
	
	fmt.Println("Digite um número inteiro de 3 dígitos: ")
	fmt.Scan(&n)

	if n > 99 && n < 1000 {
		a = n / 100
	    b = (n - a*100) / 10
		fmt.Println("O número da dezena é: ", b)		
	} else {
		fmt.Println("O número não tem 3 dígitos.")
	}
	
}
   
