package main 
import "fmt"

func main() {
 var numeros1 [10]int
 var numeros2 [5]int
 for i := 0; i < 10; i++ {
  fmt.Printf("Digite o %dº número para o primeiro vetor: ", i+1)
  fmt.Scan(&numeros1[i])
 }
 for i := 0; i < 5; i++ {
  fmt.Printf("Digite o %dº número para o segundo vetor: ", i+1)
  fmt.Scan(&numeros2[i])
 }
 var soma1 = 0
 var soma2 = 0
 for i := 0; i < 10; i++ {
  if numeros1[i] % 2 == 0 {
   soma1 = soma1 + numeros1[i]
  }
 }
 for i := 0; i < 5; i++ {
 soma2 = soma2 + numeros2[i]
 }
 
var soma3 = 0
for i := 0; i < 10; i++ {
	if numeros1[i] % 2 != 0 {
		soma3 = soma3 + numeros1[i]
	}
}
	fmt.Printf("A soma dos números pares do primeiro vetor mais a soma do segundo vetor é: %d\n", soma1+soma2)
	fmt.Printf("A soma dos números ímpares do primeiro vetor mais a soma do segundo vetor é: %d\n", soma2+ soma3)

}
 