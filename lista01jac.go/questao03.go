package main

import "fmt"

func main() {
	var
	n1, n2, n3, resultado, quadrado int
 
	 if n1<0 || n1>9 || n2<0 || n2>9 || n3<0 || n3>9 {	
		fmt.Println("As notas devem ser entre 0 e 9.")
		return
	}
	fmt.Print("Digite os numeros : ")
	fmt.Scan(&n1, &n2, &n3)
	
	resultado = n1*100 + n2*10 + n3	
	
	quadrado = resultado * resultado
    
	fmt.Printf("O número formado é %d, e o quadrado é %d", resultado, quadrado)
	
}