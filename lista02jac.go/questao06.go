package main

import f "fmt"

func main() {

	var numero1, numero2 int
	
	f.Print("Digite o primeiro número:")
	f.Scan(&numero1)
	f.Print("Digite o segundo número:")
	f.Scan(&numero2)
	if  numero1 % numero2 == 0{
		f.Printf("A é divisível por B")
	} else {
		f.Printf("A não é divisível por B")
	}
		
}