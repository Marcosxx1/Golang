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
	variavel_string := "Frase em string"
	variavel_int := 10
	variavel_double := 1.5
	variavel_boolean := true

	fmt.Printf("%s é do tipo %T\n", variavel_string, variavel_string)
	fmt.Printf("%d é do tipo %T\n", variavel_int, variavel_int)
	fmt.Printf("%f é do tipo %T\n", variavel_double, variavel_double)
	fmt.Printf("%t é do tipo %T\n\n", variavel_boolean, variavel_boolean)

	var nota int
	nota = 10

	var texto string
	texto = "Sua nota é: "

	// Criação e atribuição
	// Quando criamos e atribuimos, Go entende que a variável será daquele tipo
	mensagem := "texto criado e inicializado!"

	fmt.Printf("%s %d", texto, nota)
	fmt.Printf("\n%s\n ", mensagem)

	// tipos de dados
	TextoAlterado := "teste"
	TextoAlterado = "novo valor da variável texto"

	fmt.Println(TextoAlterado)
}
