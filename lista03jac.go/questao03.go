package main
import "fmt"
func main() {
  const numNotas int = 10 
  var (
	nota [numNotas]float64 
  )
	for i := 0; i < numNotas; i++ {
		fmt.Printf("Informe a nota %d: ", i)
		fmt.Scan(&nota[i])
	}
	for i := numNotas - 1; i >= 0; i-- {
		fmt.Printf("Nota %d = %.0f\n", i, nota[i])
	}
}