package main
import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

  if numero % 3 == 0 && numero % 5 == 0 {
	 fmt.Print("O número é divisivel por 3 e 5.")
  } else {
	 fmt.Print("O número não é divisivel por 3 e 5.")
  }
}