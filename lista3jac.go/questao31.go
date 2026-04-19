package main 
import "fmt"
func main(){
	var soma uint64 = 0
	var n uint64 = 1
	for i:= 1; i<= 64; i++ {
		soma = soma + n 
		n = n * 2
		
	}
fmt.Printf("O número de grãos será %d",soma)
}