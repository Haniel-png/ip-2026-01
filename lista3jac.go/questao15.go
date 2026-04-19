package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite um número positivo:")
	fmt.Scan(&n)
 for i := 1; i <= n; i++ {
   termo := i * i
   fmt.Printf("%d,", termo)
 }
}
