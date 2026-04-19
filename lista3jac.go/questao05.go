package main
import "fmt"
func main() {
   var idade, continuar int
   var altura, peso float64
   var TotalPessoas, qtdMais50, qtd10_20, qtd40 int
   var somaAltura10_20 float64

  fmt.Print("---Sistema de cadastro de pessoas---\n")
   
  for {
	    fmt.Print("Digite a idade da pessoa: ")
		fmt.Scan(&idade)
		
		fmt.Print("Digite a altura da pessoa: ")
		fmt.Scan(&altura)
		
		fmt.Print("Digite o peso da pessoa: ")
		fmt.Scan(&peso)
		
		TotalPessoas++
		if idade > 50 {
			qtdMais50++
		}	
		if idade >= 10 && idade <= 20 {
			qtd10_20++
			somaAltura10_20 = somaAltura10_20 + altura
		}	
		if peso < 40 {
			qtd40++
		}
		fmt.Print("Deseja cadastrar outra pessoa? (1 para sim, Qualquer outro valor para não): ")
		fmt.Scan(&continuar)
if continuar != 1 {
	break
     }
 }  
 fmt.Print("====Resultados====\n")
 fmt.Printf("Total de pessoas cadastradas: %d\n", TotalPessoas)
 fmt.Printf("Pessoas com mais de 50 anos: %d\n", qtdMais50)
 fmt.Printf("Pessoas com idade entre 10 e 20 anos: %d\n", qtd10_20)
 if qtd10_20 > 0 {
	 fmt.Printf("Média das alturas das pessoas com idade entre 10 e 20 anos: %.2f\n", somaAltura10_20/float64(qtd10_20))
 }
 if TotalPessoas > 0 {
	 fmt.Printf("Porcentagem de pessoas com peso inferior a 40kg: %.2f%%\n", (float64(qtd40)/float64(TotalPessoas))*100)
 }
}