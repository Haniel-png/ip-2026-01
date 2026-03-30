package main

import (
	"fmt"
)

func main() {
	var a1, r, n int
	fmt.Printf("Digite o primeiro termo, a razão e o número de termos da PA: ")
	fmt.Scan(&a1, &r, &n)	
	 
	soma := 0
	 
	termo := a1
	
	for i := 1; i <= n; i++ {
		soma = soma + termo
		termo = termo + r
	 }
	fmt.Printf("Soma dos termos da PA: %d\n", soma)
}