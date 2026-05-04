package main

import "fmt"

func main() {
	var idades [50]int
	var freq [121]int 
	var moda, maiorFreq int


	fmt.Println("Digite 50 idades:")
	for i := 0; i < 50; i++ {
		fmt.Scan(&idades[i])

	
		if idades[i] >= 0 && idades[i] <= 120 {
			freq[idades[i]]++
		}
	}

	
	moda = 0
	maiorFreq = freq[0]

	for i := 1; i <= 120; i++ {
		if freq[i] > maiorFreq {
			maiorFreq = freq[i]
			moda = i
		}
	}


	fmt.Printf("\nA moda das idades é: %d\n", moda)
	fmt.Printf("Frequência: %d vezes\n", maiorFreq)
}