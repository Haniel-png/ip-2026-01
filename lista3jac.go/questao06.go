package main
import "fmt"
func main() {
   var numtriangular, produto int

   fmt.Print("Digite um número inteiro positivo: ")
   fmt.Scan(&numtriangular)
   if numtriangular < 0 {
	   fmt.Println("Por favor, digite um número inteiro positivo.")
	   return
   }
    
	x := 1 
for  x * (x+1) * (x+2) < numtriangular {
		x = x + 1
	}
	produto = x * (x+1) * (x+2)
	if produto == numtriangular {
		fmt.Printf("O número %d é um número triangular.\n", numtriangular)
	} else { fmt.Printf("O número %d não é um número triangular.\n", numtriangular) }
}
