package main
import (
	"fmt" 
	"math"
)
func main () {
	soma := 0.0
	termo := 0.0
	sinal := 1.0
	numero:= 1.0
for i := 1; i< 51; i++ {
	termo = 1.0/math.Pow(numero, 3.0) * sinal
	i = i + 2 
	numero = numero + 2.0
	sinal = sinal * -1.0
	soma = soma + termo
} 
Prod := math.Pow(soma*32, 1.0/3.0)
fmt.Printf("O valor de Pi será igual a %.2f", Prod)

}