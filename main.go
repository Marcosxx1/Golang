package main

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
}

func (e endereco) apresentar() {
	fmt.Printf("Rua: %s, numero: %d, cidade: %s", e.rua, e.numero, e.cidade)
}

func main() {
	fmt.Println("Hello World")
	instanciaEndereco := endereco{
		rua:    "Rua Meu Endereço",
		numero: 21,
		cidade: "Cidade pequena",
	}
	instanciaEndereco.apresentar()
}
