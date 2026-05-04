package main

import (
	"fmt"
	"sort"
)

type Empregado struct {
	numero int
	meses  int
}

func main() {
	var lista []Empregado

	for len(lista) < 100 {
		var num, meses int
		fmt.Print("Digite o número do empregado e meses de trabalho: ")
		fmt.Scan(&num, &meses)

		if num == 0 && meses == 0 {
			break
		}

		lista = append(lista, Empregado{num, meses})
	}

	
	sort.Slice(lista, func(i, j int) bool {
		return lista[i].meses < lista[j].meses
	})

	
	limite := 3
	if len(lista) < 3 {
		limite = len(lista)
	}

	
	fmt.Println("\nOs empregados mais recentes são:")
	for i := 0; i < limite; i++ {
		fmt.Printf("Empregado %d - %d meses de trabalho\n",
			lista[i].numero, lista[i].meses)
	}
}