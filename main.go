package main

import (
	"Golang/model"
	"fmt"
)
 
func main() {
  arquivoJson := model.Endereco{
    Rua:    "Rua Meu Endereço",
    Numero: 21,
    Cidade: "Cidade pequena",
  }
 
  pessoa := model.Pessoa{
    Nome:     "João",
    Idade:    30,
    Endereco: arquivoJson,
  }
  fmt.Println(pessoa)
  arquivoJson.Apresentar()
}