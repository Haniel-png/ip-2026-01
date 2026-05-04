package main

import "fmt"


func buscarConta(codigos [10]int, codigo int) int {
	for i := 0; i < 10; i++ {
		if codigos[i] == codigo {
			return i
		}
	}
	return -1 
}

func main() {
	var codigos [10]int
	var saldos [10]float64


	fmt.Println("Cadastro das contas:")

	for i := 0; i < 10; i++ {
		for {
			fmt.Printf("Digite o código da conta %d: ", i+1)
			fmt.Scan(&codigos[i])


			repetido := false
			for j := 0; j < i; j++ {
				if codigos[j] == codigos[i] {
					repetido = true
					break
				}
			}

			if repetido {
				fmt.Println("Código já existente! Digite outro.")
			} else {
				break
			}
		}

		fmt.Printf("Digite o saldo da conta %d: ", i+1)
		fmt.Scan(&saldos[i])
	}

	
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Depósito")
		fmt.Println("2. Saque")
		fmt.Println("3. Ativo bancário")
		fmt.Println("4. Finalizar")

		var opcao int
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		if opcao == 4 {
			fmt.Println("Programa encerrado.")
			break
		}

		switch opcao {

		case 1: 
			var codigo int
			var valor float64

			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)

			pos := buscarConta(codigos, codigo)

			if pos == -1 {
				fmt.Println("Conta não encontrada.")
			} else {
				fmt.Print("Digite o valor do depósito: ")
				fmt.Scan(&valor)
				saldos[pos] += valor
				fmt.Println("Depósito realizado.")
			}

		case 2: 
			var codigo int
			var valor float64

			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&codigo)

			pos := buscarConta(codigos, codigo)

			if pos == -1 {
				fmt.Println("Conta não encontrada.")
			} else {
				fmt.Print("Digite o valor do saque: ")
				fmt.Scan(&valor)

				if saldos[pos] >= valor {
					saldos[pos] -= valor
					fmt.Println("Saque realizado.")
				} else {
					fmt.Println("Saldo insuficiente.")
				}
			}

		case 3: 
			var total float64 = 0

			for i := 0; i < 10; i++ {
				total += saldos[i]
			}

			fmt.Printf("Ativo bancário total: %.2f\n", total)

		default:
			fmt.Println("Opção inválida.")
		}
	}
}