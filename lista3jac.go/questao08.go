package main
import "fmt"
func main() {
	var soma int

	for i := 1; i <= 20; i++ {
		soma = soma + i
	}
	fmt.Printf("A soma dos números digitados é: %d\n", soma)
}

