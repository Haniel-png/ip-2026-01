package main

import "fmt"

func main() {
	var c, f float64
	fmt.Printf("Digite a temperatura em Fahrenheit:")
	fmt.Scan(&f)
	c = (f - 32) * 5 / 9
	fmt.Printf("A temperatura em Celsius é: %.2f", c)
}
