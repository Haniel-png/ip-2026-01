package main

import (
	"fmt"
)

func main() {
	var matricula int
	var horasExtras float64

	const salarioMinimo = 788.00
	const valorHoraExtra = 10.00

	var salarioHE, salarioBruto float64
	var descontoINSS, descontoIR float64
	var salarioLiquido float64

	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scan(&matricula)

	fmt.Print("Digite a quantidade de horas-extras: ")
	fmt.Scan(&horasExtras)

	salarioHE = horasExtras * valorHoraExtra
	salarioBruto = 3*salarioMinimo + salarioHE

	if salarioBruto > 1500 {
		descontoINSS = salarioBruto * 0.12
	}

	if salarioBruto > 2000 {
		descontoIR = salarioBruto * 0.20
	}

	salarioLiquido = salarioBruto - (descontoINSS + descontoIR)

	fmt.Println("\n--- Folha de Pagamento ---")
	fmt.Printf("Matrícula: %d\n", matricula)
	fmt.Printf("Salário bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Desconto INSS: R$ %.2f\n", descontoINSS)
	fmt.Printf("Desconto IR: R$ %.2f\n", descontoIR)
	fmt.Printf("Salário líquido: R$ %.2f\n", salarioLiquido)
}