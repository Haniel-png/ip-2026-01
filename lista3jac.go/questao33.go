package main 
import "fmt"
func main (){
	var n1, n2 int
	quociente := 0
	
	fmt.Print("Digite o primeiro termo:")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo termo:")
	fmt.Scan(&n2)
	resto:= n1
	for resto >= n2 {
		resto = resto - n2
		quociente = quociente + 1 
	}

	fmt.Printf("O quociente será %d\n", quociente)
	fmt.Printf("O resto é %d\n", resto)
}