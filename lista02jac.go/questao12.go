package main

import "fmt"

func main() {
	
	var idade int
	var categoria string
	
	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)

	if 0 <= idade && idade <= 2 {
		categoria = "Recém-nascido"
	}
	if 3 <= idade && idade <= 11 {	
		categoria = "Criança"
	}
	if 12 <= idade && idade <= 19 {
		categoria = "Adolescente"
	}
	if 20 <= idade && idade <= 55 {
		categoria = "Adulto"
	}
	if idade > 55 {
		categoria = "Idoso"
	}
	fmt.Printf("A categoria é: %s", categoria)
}