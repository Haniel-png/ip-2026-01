package main

import (
	"fmt"
)

func main() {
	for i := 0; i< 10; i++ {
	for j := 0; j< 10; j++ {
			if j<i {
				fmt.Printf("[%d,%d]",i , j)
			} else{
				fmt.Print("[0,0]")
			}
		}
		fmt.Println()
	} 
}