package main

import f "fmt"

func main() {

	var numero int
	f.Print("Digite um número:")
	f.Scan(&numero)
	if numero > 0 {
		f.Printf("O número é positivo")
	} else if numero < 0 {
		f.Printf("O número é negativo")
	} else {
		f.Printf("O número é zero")
	}
}