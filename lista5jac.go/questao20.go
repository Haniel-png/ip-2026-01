package main

import "fmt"

func main() {
	var dado [20]int
	var freq [20]int 
	fmt.Println("Digite 20 números inteiros (1 a 6):")
	for i := 0; i < 20; i++ {
		fmt.Scan(&dado[i])
		freq[dado[i]]++
	}
	fmt.Println("Frequência de cada número:")
	for i := 1; i <= 6; i++ {
		fmt.Printf("Número %d: %d vezes\n", i, freq[i])
	}
}

