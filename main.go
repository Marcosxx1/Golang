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
    Idade:    0,
		DataNascimento: time.Date(2000, 11, 25, 0,0,0,0, time.Local),
    Endereco: arquivoJson,
  }


  fmt.Println("pessoa: ", pessoa)
	pessoa.CalculaIdade()
  fmt.Println("idade: ", pessoa.Idade)
  arquivoJson.Apresentar()
	

}