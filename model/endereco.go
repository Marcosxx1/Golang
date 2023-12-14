package model

import (
	"encoding/json"
	"fmt"
)

/*
	Encapsulamento em Golang:
	Letras maiúsculas são públicas,
	Letras minúsculas são privadas.
	Podemos ter métodos públicos e privados.
	Podemos ter campos públicos e privados.
*/
type Endereco struct {
  Rua    string `json:"rua"`
  Numero int    `json:"numero"`
  Cidade string `json:"cidade"`
  codigo int
}
 
func (e Endereco) Apresentar() { // método público 
  jsonData, err := json.Marshal(e)
  if err != nil {
    fmt.Println("Erro ao converter o endereço para JSON:", err)
    return
  }
 
  fmt.Print(string(jsonData))
}
/* Criar um método privado que salva o nome do endereço */

func (e *Endereco) setCodigo(codigo int) {
	e.codigo = 1
}