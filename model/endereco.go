package model

import (
	"encoding/json"
	"fmt"
)

type Endereco struct {
  Rua    string `json:"rua"`
  Numero int    `json:"numero"`
  Cidade string `json:"cidade"`
  codigo int
}
 
func (e Endereco) Apresentar() {
  jsonData, err := json.Marshal(e)
  if err != nil {
    fmt.Println("Erro ao converter o endereço para JSON:", err)
    return
  }
 
  fmt.Print(string(jsonData))
}