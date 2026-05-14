package main

import "fmt"

type Criatura struct {
	Nome string
	Peso float64
	Apelido string
}

func main() {
	c := Criatura{
		Nome: "Braun",
		Peso: 120.5,
		Apelido: "Viadinho",
	}
	fmt.Println(c.Nome)
	fmt.Println(c.Peso)
	fmt.Printf("Apelido %s\n", c.Apelido)
}

