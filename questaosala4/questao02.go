package main
import "fmt"
func BuscaBinaria(array []int, x int)int {
	e:= 0
	d:=len(array)-1
	for e <= d {
		m := (e + d)/2
		if array[m]== x {
			return m 
		}
		if array[m] < x {
			e = m + 1
		}
		if array[m] > x {
			d = m - 1
		}
	}
	return -1
}
func main (){
	array := [] int {1, 4, 7, 12, 17, 21, 32}
	fmt.Print(BuscaBinaria(array,21))
}
