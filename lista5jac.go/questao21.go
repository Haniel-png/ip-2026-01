package main 
import "fmt"
func main () {
	var vetor[10]float64
	var codigo int
	fmt.Print("Digite o código ( 0, 1 ou 2): ")
	fmt.Scan(&codigo)
	fmt.Println("Digite 10 números:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&vetor[i])
	}
switch codigo {
	case 0:
		fmt.Println("Programa encerrado.")
		return
	case 1:
		fmt.Println("\nNúmeros na ordem direta:")
		for i := 0; i < 10; i++ {
			fmt.Print(vetor[i], " ")
		}
		fmt.Println()
	case 2:
		fmt.Println("\nNúmeros na ordem inversa:")
		for i := 9; i >= 0; i-- {
			fmt.Print(vetor[i], " ")
		}
		fmt.Println()
}
}