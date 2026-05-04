package main
import "fmt"
func main() {
	var fib [50]int
	fib [0] = 0
	fib [1] = 1

	for i := 2; i < 50; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Println("Os 50 primeiros números da sequência de Fibonacci são:")
	for i := 0; i < 50; i++ {
		fmt.Printf("%d ", fib[i])
		fmt.Println()
	}
}