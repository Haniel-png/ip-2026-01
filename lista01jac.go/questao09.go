package main
import "fmt"

func main() {
	var A, B, C float64
	var delta float64

	fmt.Printf("Digite os coeficientes A, B e C da equação do segundo grau:")
	fmt.Scan(&A, &B, &C)
	delta = B*B - 4*A*C

	fmt.Printf("O valor de delta é: %.2f\n", delta)
}