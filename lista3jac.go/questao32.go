package main 
import "fmt"
func main (){
	var N1, N2 int
	soma := 0
	fmt.Print("Digite o primeiro termo:")
	fmt.Scan(&N1)
	fmt.Print("Digite o segundo termo:")
	fmt.Scan(&N2)
	for i:= 1; i<=N1; i++ {
		soma = soma + N2
	}
	fmt.Printf("O valor de N1 * N2 será igual a %d",soma)
}