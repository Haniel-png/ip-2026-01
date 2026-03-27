package main

import f "fmt"

func main() {

	var numero int

	f.Print("Digite um número inteiro: ")
	f.Scan(&numero)
	if numero%5 == 0 {
		f.Print("O número é divisível por 5.")
	} else {
		f.Print("O número não é divisível por 5.")
	}
}