package main

import (
	"fmt"
	"sort"
)

func main() {
	var v [10]int

	
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}

	
	sort.Ints(v[:]) 

	
	fmt.Println("\nVetor em ordem crescente:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", v[i])
	}
	fmt.Println()
}