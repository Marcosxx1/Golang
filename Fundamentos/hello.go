package main

import (
	"fmt"
)

func main() {
	// var nome_variavel tipo_variavel
	// Quando criamos uma variavel, ela é inicializada de forma automática, por exemplo:
	// variavel_iniciada string
	// fmt.Println(variavel_iniciada)
	// a saída será um campo vazio, pois não ha valor na variável
	// o mesmo serve para qualquer variável
	var nome = "marcos"
	var peso = 10
	fmt.Printf("Hello, %s e o peso é %d\n", nome, peso)
	fmt.Printf("O tipo da variável nome é %T e o tipo da variável peso é %T\n", nome, peso)

	// criação e atribuição

	texto := "Este é um texto"
	numero := 10.1

	fmt.Printf("%s, e o valor é: %.2f", texto, numero)
}
