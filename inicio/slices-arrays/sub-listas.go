package slicesarrays

import "fmt"

func subListas() {
	// Criando listas a partir de uma existente
	var listaTotal = []int{2, 10, 9, 11, 23, 34, 3, 4, 5}
// pega os [:n], onde n é 3.
segundaLista := listaTotal[:3]
// A segundaLista inclui os elementos do início até o índice 2 (3 - 1).

// pega os [n:], onde n é 3.
terceiraLista := listaTotal[3:]
// A terceiraLista inclui os elementos do índice 3 até o final da lista.

// pega os [n:m], onde n é 2 e m é 4.
quartaLista := listaTotal[2:4]
// A quartaLista inclui os elementos do índice 2 até o índice 3 (4 - 1).

	

	fmt.Println(listaTotal)
	fmt.Println(segundaLista)
	fmt.Println(terceiraLista)
	fmt.Println(quartaLista)


	/* ARRAYS */
	// criando um array
	var arrayTotal [10]int
	// preenchendo o array
	for i := 0; i < len(arrayTotal); i++ {
		arrayTotal[i] = i
	}
	// imprimindo o array
	fmt.Println(arrayTotal)

	/* A diferença entre um array e um slice é que um slice é uma referência a um array.
	Quando você altera um slice, você altera o array original.
	Quando você altera um array, você altera todos os slices que apontam para esse array.
	Se você alterar um array, você pode alterar todos os slices que apontam para esse array.
	Se você alterar um slice, você pode alterar o array original.
	 */
}
