package main
import "fmt"
func main() {
	var x float64
	
	fmt.Print("Digite um número real:")
	fmt.Scan(&x)
	var soma float64 = 0
	var fatorial float64 = 1
	var sinal float64 = 1
	for i := 0 ; i < 20; i++ {
		if i > 0 {
		fatorial = fatorial * float64(i)
	}
	termo := ( x / fatorial) * sinal
	soma = soma + termo
	sinal = sinal * -1.0
}
fmt.Print("O somatório dos 20 primeiros termos é:", soma)
}