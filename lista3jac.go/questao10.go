package main
import "fmt"
func main() {
	var N int
	fmt.Print("Digite a quantidade de termos da sequência de Fibonacci(N>=3): ")
	fmt.Scan(&N)
	if N < 3 {
		fmt.Println("Por favor, digite um número inteiro maior ou igual a 3.")
		return
	}
	a, b := 0, 1
	fmt.Printf("Sequência de Fibonacci com %d termos:\n", N)
	fmt.Printf("%d - %d ", a, b)
	for i := 3 ; i <= N ; i++ {
		c := a + b
		fmt.Printf("- %d ", c)
		a = b
		b = c
}
	fmt.Println()
}