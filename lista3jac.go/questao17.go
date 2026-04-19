package main

import (
	"fmt"
)

func main() {
	fmt.Print("Matrix 10x10\n")
	for i:= 0; i< 10; i++ {
		for j:= 0; j< 10; j++ {
		fmt.Printf("[%d,%d]\t", i, j)
		}
		fmt.Println()
	}
}