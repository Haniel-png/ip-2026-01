package main

import f "fmt"

func main() {

	var numero int

	f.Print("Digite um número:")
	f.Scan(&numero)

	if numero%2 == 0 {
		f.Printf("É par")
	} else {
		f.Printf("É ímpar")
	}

}

