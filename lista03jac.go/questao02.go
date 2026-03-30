package main
import "fmt"
func main() {
  const numNotas int = 5 
  var (
	nota [numNotas]float64 
	soma float64 = 0
	  )
	  fmt.Printf("Informe a nota 1: ")
	  fmt.Scan(&nota[0]) 
	  fmt.Printf("Informe a nota 2: ")
	  fmt.Scan(&nota[1]) 
	  fmt.Printf("Informe a nota 3: ")
	  fmt.Scan(&nota[2])
	  fmt.Printf("Informe a nota 4: ")
	  fmt.Scan(&nota[3])
	  fmt.Printf("Informe a nota 5: ")
	  fmt.Scan(&nota[4])
	soma = nota[0] + nota[1] + nota[2] + nota[3] + nota[4]	
	fmt.Printf("A média é: %f\n ", soma/float64(numNotas))
}