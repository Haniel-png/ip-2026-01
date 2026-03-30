package main
import "fmt"
func main() {
  const numNotas int = 3 
  var (
	nota [numNotas]float64 
	soma float64 = 0
	  )
	fmt.Print("Informe a nota 1: ")
	fmt.Scan(&nota[0]) 
	fmt.Print("Informe a nota 2: ")
	fmt.Scan(&nota[1]) 
	fmt.Print("Informe a nota 3: ")
	fmt.Scan(&nota[2])
	soma = nota[0] + nota[1] + nota[2]

	fmt.Printf("A média é: %f\n ", soma/float64(numNotas))
}