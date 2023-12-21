package main

import "fmt"

func operadoresEConstantes() {
	num1 := 10
	num2 := 20
	result := num1 + num2
	fmt.Println(result)

	// constantes, os valores nunca são alterados
	// não usamos :=, MAS apenas =
	const idCliente = 10
	fmt.Printf("A id do cliente é: %d e o tipo é %T", idCliente, idCliente)
}
