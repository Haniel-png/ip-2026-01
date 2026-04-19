package main
import "fmt"
func main() {
	var n int
	var totalAprovados, totalExame, totalReprovados int
	var somaClasse float64
	
	fmt.Print("Digite o número de alunos: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var nota1, nota2 float64
		fmt.Printf("Digite a primeira nota do aluno %d: ", i)
		fmt.Scan(&nota1)
		fmt.Printf("Digite a segunda nota do aluno %d: ", i)
		fmt.Scan(&nota2)
		media := (nota1 + nota2) / 2
		somaClasse = somaClasse + media
		if media >= 7 {
			totalAprovados = totalAprovados + 1
			fmt.Print("Aluno aprovado.\n")
		} else if media >= 3 && media < 7 {
			totalExame = totalExame + 1
			fmt.Print("Aluno em exame.\n")
		} else {
			totalReprovados = totalReprovados + 1
			fmt.Print("Aluno reprovado.\n")
		}
	}
	fmt.Printf("Total de alunos aprovados: %d\n", totalAprovados)
	fmt.Printf("Total de alunos em exame: %d\n", totalExame)
	fmt.Printf("Total de alunos reprovados: %d\n", totalReprovados)
	fmt.Printf("Média da classe: %.2f\n", somaClasse/float64(n))
}