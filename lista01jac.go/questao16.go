package main

import (
	"fmt"
)

func main() {
	var salario float64

	
	fmt.Scanln(&salario)

	if salario <= 300 {
		salario = salario + salario*0.5
	} else {
		salario = salario + salario*0.3
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salario)
}