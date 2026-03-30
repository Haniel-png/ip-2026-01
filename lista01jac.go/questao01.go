package main

import "fmt"

func main() {
	var 
	  n1, n2, n3 float64

	fmt.Print("Digite a primeira nota: ")
	fmt.Scanln(&n1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scanln(&n2)
	fmt.Print("Digite a terceira nota: ")
	fmt.Scanln(&n3)		
	
	media := (n1 + n2 + n3) / 3

	if media >= 6	{
		fmt.Printf("Aprovado! Média: %.2f", media)
	} else {
		fmt.Printf("Reprovado! Média: %.2f", media)
	}
	 
}

