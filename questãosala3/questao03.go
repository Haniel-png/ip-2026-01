package main

import "fmt"

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	 resultado := 1
	for i := 1; i <= n; i++ {
		resultado = resultado * i
	}
	return resultado
}
func main() {
	var n int
	fmt.Print("Digite um número: ")	
	fmt.Scan(&n)
	resultado := fatorial(n)
	fmt.Printf("O fatorial de %d é: %d\n", n, resultado)
}