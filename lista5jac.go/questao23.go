package main

import "fmt"

func main() {
	var janela [24]int   
	var corredor [24]int 

	for {
		var opcao int

		fmt.Println("\n--- Venda de Passagens ---")
		fmt.Println("1 - Janela")
		fmt.Println("2 - Corredor")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha: ")
		fmt.Scan(&opcao)

		if opcao == 3 {
			fmt.Println("Encerrando...")
			break
		}

		var vetor *[24]int

		if opcao == 1 {
			vetor = &janela
			fmt.Println("\nPoltronas disponíveis na JANELA:")
		} else if opcao == 2 {
			vetor = &corredor
			fmt.Println("\nPoltronas disponíveis no CORREDOR:")
		} else {
			fmt.Println("Opção inválida.")
			continue
		}

		
		disponivel := false
		for i := 0; i < 24; i++ {
			if vetor[i] == 0 {
				fmt.Printf("%d ", i)
				disponivel = true
			}
		}

		if !disponivel {
			fmt.Println("\nNenhuma poltrona disponível nesse setor.")
			continue
		}

		
		var lugar int
		fmt.Print("\nEscolha o número da poltrona: ")
		fmt.Scan(&lugar)

		if lugar < 0 || lugar >= 24 {
			fmt.Println("Número inválido.")
			continue
		}

		if vetor[lugar] == 1 {
			fmt.Println("Poltrona já ocupada.")
		} else {
			vetor[lugar] = 1
			fmt.Println("Passagem vendida com sucesso!")
		}

		
		cheio := true
		for i := 0; i < 24; i++ {
			if janela[i] == 0 || corredor[i] == 0 {
				cheio = false
				break
			}
		}

		if cheio {
			fmt.Println("\nÔnibus lotado!")
			break
		}
	}
}