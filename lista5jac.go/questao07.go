package main 
import "fmt"

func main() {
	var vetor[100]int
	var x int = 1
	for i:= 0 ; i < 100; i++ {
		vetor[i] = x
		x = x + 2
	}
	fmt.Println("Os 100 primeiros impares")
	for i := 0; i < 100; i++ {
		fmt.Printf("O número %d \n", vetor[i])
	}
}
