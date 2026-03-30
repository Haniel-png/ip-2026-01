package main

import "fmt"

func main() {
	var c, f float64
	var p, mm float64
	fmt.Printf("Digite a temperatura em Fahrenheit:")
	fmt.Scan(&f)
	fmt.Printf("Digite a quantidade de chuva em polegadas:")
	fmt.Scan(&p)

	c = (f - 32) * 5 / 9
	mm = p * 25.4

	fmt.Printf("A temperatura em Celsius é: %.2f\n", c)
	fmt.Printf("A quantidade de chuva em milímetros é: %.2f", mm)
}
