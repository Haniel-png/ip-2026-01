package main 
import "fmt"
func main () {
	var numero int
	var peso float64

	var numGordo, numMagro int
	var pesoGordo, pesoMagro float64
	fmt.Scan(&numero, &peso)
	pesoGordo = peso
	pesoMagro = peso
    numGordo = numero
	numMagro = numero
	for i := 2; i<=4; i++{
		fmt.Scan(&numero, &peso)
		if pesoGordo<peso {
			pesoGordo = peso
			numGordo = numero
		}
		if pesoMagro > peso{
			pesoMagro = peso
			numMagro = numero
		}
	}
	fmt.Printf("O boi mais pesado tem numero %d e peso %f\n",numGordo, pesoGordo)
    fmt.Printf("O boi mais leve tem numero %d e peso %f\n ",numMagro, pesoMagro)
}