package main 
import ("fmt"
        "math")

func main() {
	var vetor[100]float64
	var soma float64 = 0
	for i := 0; i < 100; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
	}
	for i:= 0 ; i < 50; i++ {
		termo := math.Pow(float64(vetor[i] - vetor[99-i]), 3)
		soma = soma + termo
	}

	fmt.Printf("A soma dos cubos das diferenças é: %f\n", soma)
}
