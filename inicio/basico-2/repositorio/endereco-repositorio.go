package repositorio

import "Golang/inicio/basico-2/model"

type EnderecoRepositorio struct {
	Enderecos []*model.Endereco
}

func (repo *EnderecoRepositorio) FindByID(id int) *model.Endereco {
	for _, endereco := range repo.Enderecos {
		if endereco.Id == id {
			return endereco
		}
	}
	return nil
}

func (repo *EnderecoRepositorio) FindByRua(rua string) *model.Endereco {
	for _, endereco := range repo.Enderecos {
		if endereco.Rua == rua {
			return endereco
		}
	}
	return nil
}

func (repo *EnderecoRepositorio) FindByCidade(cidade string) *model.Endereco {
	for _, endereco := range repo.Enderecos {
		if endereco.Cidade == cidade {
			return endereco
		}
	}
	return nil
}
