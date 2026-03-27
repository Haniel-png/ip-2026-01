package main

import f "fmt"

func main() {

	var n1, n2, soma int
	
	f.Print("Digite o primeiro número:")
	f.Scan(&n1)
	f.Print("Digite o segundo número:")
	f.Scan(&n2)
	
	soma = n1 + n2

	if n1 + n2 > 20 {
    soma = soma + 8
		f.Printf("A soma dos números é: %d", soma)
	} else {
		soma = soma - 5
		f.Printf("A soma dos números é: %d", soma)
	}

}