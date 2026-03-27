package main

import "fmt"

func main() {

	var (
		n1, n2, p1, p2, media float64
	)

	fmt.Println("Escreva duas notas, n1 e n2 respectivamente")
	fmt.Scan(&n1, &n2)

	fmt.Println("Escreva os pesos p1 e p2, respectivamente")
	fmt.Scan(&p1, &p2)

	media = ((n1 * p1) + (n2 * p2)) / (p1 + p2)

	fmt.Println("Escreva a média ponderada ", media)
}
