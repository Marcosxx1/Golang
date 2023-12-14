package model

import "time"

type Pessoa struct {
	Nome           string
	Idade          int
	DataNascimento time.Time
	Endereco       Endereco
}

/* Funções podem estar dentro de structs, outras funções podem ser publicas ou privadas
   Métodos devem estar dentro de structs */

// Método publico
func (p Pessoa) CalculaIdade() int { 
	return time.Now().Year() - p.DataNascimento.Year()
}

//função publica 
func CalculaIdade(dataNascimento time.Time) int {
	return time.Now().Year() - dataNascimento.Year()
}

