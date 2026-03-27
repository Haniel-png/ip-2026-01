package main

import f "fmt"

func main() {

	var valor, valorvenda float64

	f.Print("Digite um valor: ")
	f.Scan(&valor)
	if valor <10 && valor > 0 {
		valorvenda = valor * 1.70
		f.Printf("O valor de venda é: ", valorvenda)
	} else if valor >= 10 && valor < 30 {
		valorvenda = valor * 1.50
		f.Printf("O valor de venda é: ", valorvenda)
	} else if valor >= 30 && valor < 50 {
		valorvenda = valor * 1.40
		f.Printf("O valor de venda é: ", valorvenda)
	} else {
		valorvenda = valor * 1.30
		f.Printf("O valor de venda é: ", valorvenda)
      }
}