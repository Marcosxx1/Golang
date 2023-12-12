package slicesarrays

import (
	"fmt"
)

func Slices() {

	/*
		Slices são estruturas de dados que permitem acessar e manipular elementos de uma lista
		slices são como arrays, mas podem ter seu tamanho definido de forma dinâmica
		São criados utilizando a função 'make'
		São referências de tipo

		Criando slices usando a função 'make':
		make([]tipo_da_variável, tamanho, capacidade)
		tamanho é o numero de elemento do slice
		capacidade é o numero de elementos que o slice pode conter
		se a capacidade não é definida, ela será definida pelo tamanho
		Se o tamanho é maior que a capacidade, o slice será rejustado
		se o tamanho é menor que a capacidade, o slice será truncado
		se o tamanho é igual a capacidade, o slice será inicializado

	*/

	// Criando uma slice de inteiros
	lista := []int{1, 2, 3, 4, 5}
	fmt.Println(lista)

	// Iterando sobre a slice usando um loop tradicional
	for i := 0; i < len(lista); i++ {
		fmt.Println(lista[i])
	}

	// média de números
	var listaDeNumeros = make([]int, 0)

	for _, v := range listaDeNumeros {
		fmt.Printf("%d", v)
	}

	listaDeNumeros = append(listaDeNumeros, 1, 2, 3, 4)
	fmt.Println(listaDeNumeros)

	var media int
	for i := 0; i < len(listaDeNumeros); i++ {
		media += (listaDeNumeros[i])

	}

	fmt.Println(media / len(listaDeNumeros))

	// Iterando sobre a slice usando for range
	for _, v := range lista {
		fmt.Printf("%d", v)
	}

	// Criando um slice com tamanho 5 e capacidade 10 usando make
	slice := make([]int, 5, 10)

	fmt.Println(slice)

	fmt.Println("Length:", len(slice))   // inicialmente teremos 5 elementos
	fmt.Println("Capacity:", cap(slice)) // inicialmente teremos 10 elementos

	// Adicionando elementos ao slice usando append
	slice = append(slice, 1, 2, 3, 4, 5, 6)
	fmt.Println(slice)

	fmt.Println("Length:", len(slice))   // Passamos a ter 8 elementos
	fmt.Println("Capacity:", cap(slice)) // Passamos a ter 20 elementos (GERALMENTE É O DOBRO DA INICIAL)

	// Criando uma slice de strings
	listaDeString := []string{"item 1", "item 2", "item 3"}

	fmt.Println("Printando com o for range")

	// Iterando sobre a slice de strings usando for range
	for i, v := range listaDeString {
		fmt.Println(i, v)
	}


	// Escolhendo numeros
	var listaTotal = []int{2, 10, 9, 11, 23, 34, 3, 4, 5}
	var listaComMenoresQueCinco = make([]int, 0)

	for i := 0; i < len(listaTotal); i++ {
		if listaTotal[i] <= 5 {
			listaComMenoresQueCinco = append(listaComMenoresQueCinco, listaTotal[i])
		}
	}
	fmt.Println(listaTotal)
	fmt.Println(listaComMenoresQueCinco)
}
