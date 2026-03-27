package main

import "fmt"

func main() {
    var x float64
    var fx float64

    fmt.Print("Digite o valor de x: ")
    fmt.Scan(&x)

    if x == 2 {
        fmt.Println("Erro: divisao por zero nao permitida (x nao pode ser 2).")
    } else {
        fx = 8 / (2 - x)
        fmt.Printf("O valor de f(x) =", fx)
    }
}