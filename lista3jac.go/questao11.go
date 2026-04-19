package main
import "fmt"
func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Por favor, digite um número inteiro positivo.")
		return
	}
	fatorial := 1
	for i := 1; i <= n; i++ {
		fatorial = fatorial * i
	}
	fmt.Printf("O fatorial de %d é: %d\n", n, fatorial)
}