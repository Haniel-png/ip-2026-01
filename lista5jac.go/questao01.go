package main 
import "fmt"

func main() {
 var numeros [10]int
 encontrou := false

 for i := 0; i < 10; i++ {
  fmt.Printf("Digite o %dº número: ", i+1)
  fmt.Scan(&numeros[i])
 }
 for i := 0; i < 10; i++ {
	if numeros[i] > 50 {
		encontrou = true
		fmt.Printf("O número %d é maior que 50.\n", numeros[i])
	} 
 }
	if !encontrou {
		fmt.Println("Nenhum número é maior que 50.")
	}
}
