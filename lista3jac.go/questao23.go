package main
import "fmt"
func main() {
	var numerador, denominador, sinal, termo float64
	numerador = 1001.0
	denominador = 0.0
	sinal = -1.0
	soma := 0.0
	fmt.Print("Soma da Série\n")
	for i:= 1; i <=333; i++ {
    numerador = numerador - 3
	denominador = denominador + 1
	sinal = sinal * -1.0
	termo = (numerador/denominador) * sinal
    soma = soma + termo
	}
	fmt.Printf("%f",soma)
}