package interfaces

import "Golang/model"

type Buscavel interface {
	FindByID(id int) []model.Endereco
	FindByRua(rua string) []model.Endereco
	FindByCidade(cidade string) []model.Endereco
}
