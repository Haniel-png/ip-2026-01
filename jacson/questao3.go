package main

import "fmt"

func main() {

	var (
		nota, x, qtd, media float64
	)

	fmt.Println("Escreva a quantidade de notas: ")
	fmt.Scan(&qtd)

	for i := 1; i <= int(qtd); i++ {
		fmt.Println("Escreva a nota dos alunos: ", i)
		fmt.Scan(&nota)
		x = x + nota
	}
	media = x / qtd
	fmt.Println("A média das notas é: ", media)
}
