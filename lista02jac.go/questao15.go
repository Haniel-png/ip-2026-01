package main

import (
	"fmt"
)

func main() {
	var dia, mes, ano int
	var m string

	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)
	fmt.Print("Digite o mês: ")
	fmt.Scan(&mes)
	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

  if mes == 1 {
		m = "Janeiro"
  }		else if mes == 2 {
		m = "Fevereiro"
  }	else if mes == 3 {		
		m = "Março"
  }	else if mes == 4 {
		m = "Abril"
  }	else if mes == 5 {			
		m = "Maio"
  }	else if mes == 6 {
		m = "Junho"
  }	else if mes == 7 {	
		m = "Julho"		
	  }	else if mes == 8 {		
		m = "Agosto"
  }	else if mes == 9 {
		m = "Setembro"
  }	else if mes == 10 {
		m = "Outubro"
  } 	else if mes == 11 {
		m = "Novembro"
  }	else if mes == 12 {
		m = "Dezembro"
  }
  			fmt.Printf("A data é: %d de %s de %d", dia, m, ano)		
}						