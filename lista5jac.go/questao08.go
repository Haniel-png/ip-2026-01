package main 
import ("fmt"
        "math")

func main() {
	var vetor[15]int
	var vetorraiz[15]float64
	for i := 0; i < 15; i++ {
		fmt.Printf("Digite o %dº número: ", i+1)
		fmt.Scan(&vetor[i])
		
	}
	fmt.Println("\nAs raízes quadradas dos números digitados são:")
	for i := 0; i < 15; i++ {
		if vetor[i] >= 0 {
		vetorraiz[i] = math.Sqrt(float64(vetor[i]))	
		fmt.Printf("A raiz quadrada de %d é %.2f\n", vetor[i], vetorraiz[i])
	} else if vetor[i] < 0 {
			fmt.Printf("A raiz quadrada de %d é -1\n", vetor[i])
		}
	}
}