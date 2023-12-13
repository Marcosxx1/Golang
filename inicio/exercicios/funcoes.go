package exercicios

import "fmt"

func main() {
	ola("Olá")
	ola("Olá")
	fmt.Println(soma(1, 2))
	fmt.Println(operacoes(1, 2))
	fmt.Println(operacoesComDefinicao(1, 2))

}

// função sem retorno
func ola(mensagem string) {
	fmt.Println(mensagem)
}

// função com retorno
func ola2(mensagem string) string {
	return mensagem
}

func soma(a int, b int) int {
	return a + b
}

// função com retorno de vários tipos
func operacoes(num1 int, num2 int) (int, int, int, int) {
	soma := num1 + num2
	subtracao := num1 - num2
	divisao := num1 / num2
	multiplicacao := num1 * num2

	return soma, subtracao, divisao, multiplicacao
}

/* Podemos também definir as variáveis diretamente no retorno, como no exemplo abaixo: */
func operacoesComDefinicao(num1 int, num2 int) (soma int, subtracao int, divisao int, multiplicacao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	divisao = num1 / num2
	multiplicacao = num1 * num2

	return
}
