package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Printf("Digite dois números inteiros: ")
	fmt.Scan(&x, &y)

	if x % 2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Printf("%d ", x + (i * 2))
		} 
	} else {
		fmt.Printf("O número é ímpar.")
	}
}