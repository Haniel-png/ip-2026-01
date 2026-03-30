package main

import "fmt"

func main() {
	var conta int
	var consumo, valordaconta float64
	var consumidor string

	fmt.Printf("Digite a sua conta :")
	fmt.Scan(&conta)

	fmt.Printf("Digite o consumo de água:", "m³")
	fmt.Scan(&consumo)

	fmt.Printf("Digite o tipo de consumidor:")
	fmt.Scan(&consumidor)

	if consumidor == "R" {
		valordaconta = 5 + (0.05 * consumo)
	} else if consumidor == "C" {
		if consumo <= 80 {
			valordaconta = 500
		} else {
			valordaconta = 500 + (0.25 * (consumo - 80))
		}
	} else if consumidor == "I" {
		if consumo <= 100 {
			valordaconta = 800
		} else {
			valordaconta = 800 + 0.04*(consumo-100)
		}
	}

}
