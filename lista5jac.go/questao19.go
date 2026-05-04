package main 
import "fmt"
func main () {
	var num [10]int
	var divisores [5]int
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
	}
	for j := 0; j < 5; j++ {
		fmt.Printf("Digite o %dº divisor: ", j+1)
		fmt.Scan(&divisores[j])
	}
	fmt.Println("\nNúmeros divisíveis pelos divisores fornecidos:")
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if divisores[j] != 0 && num[i]%divisores[j] == 0 {
				fmt.Printf("%d é divisível por %d\n", num[i], divisores[j])
			}
		}
	}	
}