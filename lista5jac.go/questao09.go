package main 
import "fmt"

func main() {
	var vetor[10]float64
	var soma float64 = 0
	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
	}
	for i := 0; i < 10; i++ {
		soma = soma + vetor[i]
	}
	media := soma / 10
	fmt.Print("Os alunos com altura maior que a média são:\n")
	for i := 0; i < 10; i++ {
		if vetor[i] > media {
			fmt.Printf("O número %.2f é maior que a média %.2f.\n", vetor[i], media)
		}
	}
}