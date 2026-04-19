package main
import "fmt"
func main() {
	var b, n int
fmt.Println("Digite o valor do expoente n: ")
fmt.Scan(&n)

fmt.Println("Digite o valor da base b:")
fmt.Scan(&b)
if n <= 1 {
	fmt.Println("Digite um expoente maior que 1")
	return
}
if b < 2 {
	fmt.Println("Digite uma base maior ou igual a 2")
	return
}
produto := 1
	for i:= 1; i<= n; i++ {
		produto = produto * b
	}
	fmt.Printf("O resultado é %d", produto)
}