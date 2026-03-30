package main
import "fmt"

func main() {
	var a, b, c, d float64
	var det float64

	fmt.Printf("Digite os coeficinetes da matriz 2x2 (a, b, c, d): ")
	fmt.Scan(&a, &b, &c, &d)
	det = a*d - b*c
	fmt.Printf("O determinante da matriz é: %.2f\n", det)
}