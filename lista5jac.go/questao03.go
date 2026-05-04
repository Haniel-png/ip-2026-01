package main 
import "fmt"

func main() {
 var numeros [10]int
 var soma = 0
 var qtd = 0
 for i := 0; i < 10; i++ {
  fmt.Printf("Digite o %dº número para o primeiro vetor: ", i+1)
  fmt.Scan(&numeros[i])
 }
 for i := 0; i < 10; i++ {
	if i % 2 == 0 {
		fmt.Printf("O número %d é par.\n", numeros[i])
		soma = soma + numeros[i]
	} 
 }
		for i := 0; i < 10; i++ {
			if i % 2 != 0 {
	fmt.Printf("O número %d é ímpar.\n", numeros[i])
		qtd = qtd + 1
	}
	
}
fmt.Printf("A soma dos números pares é: %d\n", soma)
fmt.Printf("A quantidade de números ímpares é: %d\n", qtd)
}
