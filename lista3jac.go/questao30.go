package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64
	for i := 0; i<=40; i++ {
		R = 0.5 * float64(i)
		V := 4.0/3.0 * math.Pi * math.Pow(R,3.0)
	fmt.Printf("%.2f\n", V)
	}
}