package main

import "fmt"

func main () {
    var x, y, z float64
   
	fmt.Print("Digite o valor da base: ")
	fmt.Scan(&x)

	fmt.Print("Digite o valor do expoente:")
	fmt.Scan(&y)
    
	z = x * x
	for i := 1; i <= int(y-2); i++ {
		z = z * x
	}
	fmt.Printf(" O resultado é: %.2f", z)
}