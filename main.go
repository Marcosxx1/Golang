package main

import (
	"Golang/model"
	"fmt"
	"time"
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
		DataNascimento: time.Date(2000, 11, 25, 0,0,0,0, time.Local),
    Endereco: arquivoJson,
  }

	idade := pessoa.CalculaIdade()

  fmt.Println("pessoa: ", pessoa)
  fmt.Println("idade: ", idade)
  arquivoJson.Apresentar()
	

}