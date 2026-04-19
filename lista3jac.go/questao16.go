package main

import (
	"fmt"
)

func main() {
	var a1, a2, n int
	fmt.Printf("Digite o primeiro termo(a1):")
	fmt.Scan(&a1)
	fmt.Printf("Digite o segundo termo (a2):")
	fmt.Scan(&a2)
	t1 := a1
	t2 := a2
	fmt.Printf("Digite o número da série:")
	fmt.Scan(&n)
	fmt.Printf("Série de Fetuccine gerada: %d, %d ", t1, t2)
	if n<3 {
		fmt.Print("Digite um número maior que 3")
		return
	}

	for i := 3; i <= n; i++ {
		var proximo int
		if i%2 != 0{
			proximo = t1 + t2
		} else {
			proximo = t2 - t1
		}
		fmt.Printf(", %d", proximo)
		t1 = t2
		t2 = proximo
	}
}
