package main

import f "fmt"

func main() {

	var numero int
	
	f.Print("Digite um número inteiro: ")
	f.Scan(&numero)

  if  20 < numero && numero < 90 {
		f.Print("O número está entre 20 e 90.")
  } else {
		f.Print("O número não está entre 20 e 90.")
  }
}