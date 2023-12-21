package model

/*
Encapsulamento em Golang:
Letras maiúsculas são públicas,
Letras minúsculas são privadas.
Podemos ter métodos públicos e privados.
Podemos ter campos públicos e privados.
*/
type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Id     int
}

func InicializaEnderecos() []*Endereco {
	var enderecos []*Endereco

	endereco1 := &Endereco{
		Rua:    "Avenida Principal",
		Numero: 10,
		Cidade: "Metrópole",
		Id:     3,
	}

	endereco2 := &Endereco{
		Rua:    "Travessa das Flores",
		Numero: 15,
		Cidade: "Cidade Jardim",
		Id:     4,
	}

	endereco3 := &Endereco{
		Rua:    "Rua Comercial",
		Numero: 25,
		Cidade: "Centro Urbano",
		Id:     5,
	}

	endereco4 := &Endereco{
		Rua:    "Avenida dos Esportes",
		Numero: 7,
		Cidade: "Vila Esportiva",
		Id:     6,
	}

	endereco5 := &Endereco{
		Rua:    "Rua do Comércio",
		Numero: 18,
		Cidade: "Distrito Comercial",
		Id:     7,
	}

	endereco6 := &Endereco{
		Rua:    "Alameda dos Parques",
		Numero: 12,
		Cidade: "Bairro Residencial",
		Id:     8,
	}

	endereco7 := &Endereco{
		Rua:    "Avenida da Cultura",
		Numero: 30,
		Cidade: "Cidade Cultural",
		Id:     9,
	}

	endereco8 := &Endereco{
		Rua:    "Travessa da Paz",
		Numero: 5,
		Cidade: "Vila Tranquila",
		Id:     10,
	}

	endereco9 := &Endereco{
		Rua:    "Rua dos Artistas",
		Numero: 22,
		Cidade: "Cidade das Artes",
		Id:     11,
	}

	endereco10 := &Endereco{
		Rua:    "Avenida do Saber",
		Numero: 14,
		Cidade: "Centro Educacional",
		Id:     12,
	}

	enderecos = append(enderecos, endereco1, endereco2, endereco3, endereco4, endereco5, endereco6, endereco7, endereco8, endereco9, endereco10)

	return enderecos
}
