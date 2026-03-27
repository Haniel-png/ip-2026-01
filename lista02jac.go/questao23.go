package main

import (
	"fmt"
)

func main() {
	var idade int
	var categoria string
	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)
	if 0 <= idade && idade < 16 {
		categoria = "Não eleitor"
	} else if 16 <= idade && idade < 18 {
		categoria = "Eleitor facultativo"
	} else if 18 <= idade && idade < 65 {
		categoria = "Eleitor obrigatório"
	}	else {
		categoria = "Eleitor facultativo"
	}	
	fmt.Printf("A categoria é: %s", categoria)
}