package main

import "fmt"

func media(a, b, c float64) float64 {
	return (a + b + c) / 3
}	
func main() {
	var a, b, c float64
	fmt.Print("Digite 3 números: ")
	fmt.Scan(&a, &b, &c)
	resultado := media(a, b, c)
	fmt.Printf("A média dos números é: %.2f\n", resultado)
}