package main

import f "fmt"

func main() {

	var valor, modo1, modo2 int
	   
	f.Printf("Digite como fará a viagem(1 para ida, 2 para ida e volta):")
	f.Scan(&modo1)
	f.Printf("Digite o destino da viagem(1 para cidade Norte, 2 para cidade Nordeste, 3 para cidade Centro-Oeste, 4 para cidade Sul):")
	f.Scan(&modo2)

	if modo2 == 1 {
		if modo1 == 1 {
			valor = 500
		} else {
			 valor = 900
	} 
	} else if modo2 == 2 {
		if modo1 == 1 {
			valor = 350
		} else {
			valor = 650
		}
	} else if modo2 == 3 {
		if modo1 == 1 {
			valor = 350
		} else {
			valor = 600
		}
	} else {
		if modo1 == 1 {
			valor = 300
		} else {
			valor = 550
		}
	}

	f.Printf("O valor da viagem é: %d", valor)
}
