package main

import (
	"fmt"
	"strconv"
)

func Conversoes() {


	var numero int8 = 127
	fmt.Println(numero)

	/* Digamos que recebemos um valor de uma função
	e queremos converter para outro tipo mas ainda não sabemos de que tipo ela é */

	value := 1238.12 // Exemplo simples onde temos um valor sem saber o tipo

	fmt.Printf("%T : tipo da variável antes da conversão.  %f : valor \n", value, value) // Exibe o tipo da variável

	valueConvertedToFloat32 := float32(value) // Converte para float32

	fmt.Printf("%T : tipo da variável depois da conversão.  %f : valor", valueConvertedToFloat32, valueConvertedToFloat32) // Exibe o tipo da variável

	booleanValue, err := strconv.ParseBool("true") // Converte para boolean "true" ou "false"

	fmt.Println("\n booleanValue", booleanValue, err)

	// Convertendo string para numero

	preco := "100.23"

	precoConvertido, _ := strconv.ParseFloat(preco, 64)
	fmt.Printf("\n preco convertido: %f", precoConvertido)
}
