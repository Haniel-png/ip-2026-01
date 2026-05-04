package main 
import "fmt"
func main () {
	var v1, v2 [10]int
	var v3 [20]int
	

	for i:= 0; i < 10; i ++ {
		fmt.Printf("Digite o %d° do v1: ", i+1)
		fmt.Scan(&v1[i])
	}
	fmt.Println("Digite os 10 valores do vetor 2:")
	for i := 0; i < 10; i++ {
		fmt.Print("Digite o %d° do v2: ", i+1)
		fmt.Scan(&v2[i])
	}
	for i := 0; i < 10; i++ {
		v3[2*i] = v1[i]
		v3[2*i+1] = v2[i]
	}
	fmt.Println("\nVetor resultante da intercalação:")
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", v3[i])
	}
	fmt.Println()
}