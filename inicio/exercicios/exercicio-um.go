package exercicios

func Exercicio1() {

	/* Relembrando
	Arryay: 	var arrayTotal [10]int
	Slice: 	var sliceTotal []int
	*/
	/* Criar um array com duas posições de inteiros
	 e armazenar em uma variável a soma total da lista
	a variável deve ser impromida no console */

	var valorTotal = [2]int{2, 3}
	var somaTotal int

	for i := 0; i < len(valorTotal); i++ {
		somaTotal += valorTotal[i]
	}
	println(somaTotal)

	/* Dado um slice com os items {2, 8, 3, 10, 5, 4, 7, 9, 1} que vão de 1 a 10
	efetuar a soma em duas variáveis, a primeira de números de 1 a 5 e a segunda
	de 6 a 10 */

	var sliceTotal = []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	var somaDeUmAteCinco int
	var somaDeSeisAteDez int
	for i := 0; i < len(sliceTotal); i++ {
		if sliceTotal[i] <= 5 {
			somaDeUmAteCinco += sliceTotal[i]
		} else if sliceTotal[i] >= 6 {
			somaDeSeisAteDez += sliceTotal[i]
		}
	}
	println(somaDeUmAteCinco)
	println(somaDeSeisAteDez)
}
