package interfaces

import "Golang/inicio/basico-2/model"

type Buscavel interface {
	FindByID(id int) []model.Endereco
	FindByRua(rua string) []model.Endereco
	FindByCidade(cidade string) []model.Endereco
}
