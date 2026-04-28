package main 
import "fmt"
func potencia(x, n int)float64 {
	if n==0 {
		return 1
	}
	if n<0{
		return 1.0/potencia(x,-n)
	}
	return float64(x) * potencia(x, n-1)
}
func main() {
	var x, n int
	fmt.Print("Digite o valor da base:")
	fmt.Scan(&x)

	fmt.Print("Digite o valor do expoente:")
	fmt.Scan(&n)
	if x==0 && n < 0 {
		fmt.Print("Indeterminação matemática")
	}

resultado := potencia(x,n)
fmt.Printf("O valor de %x, elevado a %d é igual a %f", x, n, resultado)
}