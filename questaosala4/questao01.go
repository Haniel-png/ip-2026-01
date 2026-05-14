package main

import (
	"fmt"
)

func BuscaSequencial(array []int, x int)int {
for i:= 0; i <=len(array); i++ {
	if array[i]== x{
		return i
	}
} 
return -1
}
func main () {
	array := []int {1, 2, 5, 7, 12}
	fmt.Print(BuscaSequencial(array,5))
	
}