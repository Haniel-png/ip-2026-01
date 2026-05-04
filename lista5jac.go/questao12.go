package main

import "fmt"

func main() {
	var notas [15]int
	var freqAbs [11]int 
	var freqRel [11]float64

	for i := 0; i < 15; i++ {
		fmt.Printf("Digite a %dª nota (0 a 10): ", i+1)
		fmt.Scan(&notas[i])

		if notas[i] < 0 || notas[i] > 10 {
			fmt.Println("Nota inválida! Digite novamente.")
			i-- // volta uma posição
			continue
		}

freqAbs[notas[i]]++
	}

		for i := 0; i <= 10; i++ {
		freqRel[i] = float64(freqAbs[i]) / 15.0
	}

	fmt.Println("\nNota | Freq. Absoluta | Freq. Relativa")
	fmt.Println("--------------------------------------")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%4d | %15d | %16.2f\n", i, freqAbs[i], freqRel[i])
	}
}