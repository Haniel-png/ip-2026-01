package main

import "fmt"

func main() {
	preco := 6.0
	qtd := 130.0

	maiorLucro := -999999.0
	melhorPreco := 0.0
	melhorQtd := 0.0

	fmt.Println("Preço | Quantidade | Lucro")

	for preco >= 1.0 {
		lucro := preco*qtd - 300.0

		fmt.Printf("R$ %.2f | %.0f | R$ %.2f\n", preco, qtd, lucro)

		if lucro > maiorLucro {
			maiorLucro = lucro
			melhorPreco = preco
			melhorQtd = qtd
		}

		preco -= 0.6
		qtd += 30
	}

	fmt.Println("\n--- Resultado ---")
	fmt.Printf("Maior lucro: R$ %.2f\n", maiorLucro)
	fmt.Printf("Melhor preço: R$ %.2f\n", melhorPreco)
	fmt.Printf("Quantidade de ingressos: %.0f\n", melhorQtd)
}