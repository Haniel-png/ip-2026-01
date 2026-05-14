package main
import "fmt"
func main() {
  const numNotas int = 5 
  var (
	nota [numNotas]float64 
	soma float64 = 0
	  )
	for i := 0; i < numNotas; i++ {
	  fmt.Printf("Informe a nota %d: ", i+1)
	  fmt.Scan(&nota[i]) 
	}
	for i := 0; i < numNotas; i++ {
		soma = soma + nota[i]
	}
	fmt.Printf("A média é: %f\n ", soma/float64(numNotas))
}