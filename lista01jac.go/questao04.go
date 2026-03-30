package main

import "fmt"

func main() {
	var salario float64
	var custo float64
	var consumo float64
	var custo1 float64

	fmt.Print("Digite o seu salário: ")
	fmt.Scanln(&salario)

	fmt.Print("Digite a quantidade de KW gasto: ")
	fmt.Scanln(&consumo)

	valordecadaKW := (70 * salario) / 100

	custo := valordecadaKW * consumo

    custo := valordecadaKW*consumo

    custo1 := (custo*90)/100

fmt.Printf("Custo por KW: %.2f\n", valordecadaKW)
fmt.Printf("Custo do consumo: %.2f\n", custo)
fmt.Printf("Custo com o desconto: %.2f\n", custo1)
}