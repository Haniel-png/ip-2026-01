package main

import "fmt"

func main(){
	var soma, termo float64
	var numerador, denominador float64
	numerador = 39.0
	denominador = 0.0
for i:= 1; i <= 37; i++  {
	numerador = numerador - 1
	denominador = denominador + 1
termo = ((numerador)*(numerador - 1))/(denominador)
soma = soma + termo
}
fmt.Printf("O resutado é %f", soma)
}