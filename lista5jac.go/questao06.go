package main 
import "fmt"

func main() {
	var vetor[100]int
	for i:= 0 ; i < 100; i++ {
		vetor[i] = 100 - i
	}
	for i := 0; i < 100; i++ {
		fmt.Printf("O número %d está na posição %d do vetor.\n", vetor[i], i)
	}
}