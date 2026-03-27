package main

import f "fmt"

func main() {

	var num1, num2, num3 int
	f.Print("Digite o primeiro número:")
	f.Scanln(&num1, &num2, &num3)
	if num1 > num2 && num1 > num3 {
		if num2 > num3 {
			f.Printf("O maior número é: %d", num1)
			f.Printf("O intermediário número é: %d", num2)
			f.Printf("O menor número é: %d", num3)
		} else {
			f.Printf("O maior número é: %d", num1)
			f.Printf("O intermediário número é: %d", num3)
			f.Printf("O menor número é: %d", num2)
		}
	} else if num2 > num1 && num2 > num3 {
		if num1 > num3 {
			f.Printf("O maior número é: %d", num2)
			f.Printf("O intermediário número é: %d", num1)
			f.Printf("O menor número é: %d", num3)
		} else {
			f.Printf("O maior número é: %d", num2)
			f.Printf("O intermediário número é: %d", num3)
			f.Printf("O menor número é: %d", num1)
		}	
	} else {
		if num1 > num2 {
			f.Printf("O maior número é: %d", num3)
			f.Printf("O intermediário número é: %d", num1)
			f.Printf("O menor número é: %d", num2)
		} else {		
			f.Printf("O maior número é: %d", num3)
			f.Printf("O intermediário número é: %d", num2)
			f.Printf("O menor número é: %d", num1)
		}
	}
}