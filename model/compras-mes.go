package model

import "time"

type ComprasMes struct {
	DataCompra       time.Time
	Mercado          Mercado
	ItemsParaComprar []ItemsCompra
}

func IniciaCompraMes() *ComprasMes {

	localDeCompra := &Mercado{
		Endereco: "Rua do Mercado, 123",
		Nome:     "Mercado Mercadil",
	}

	itemsCompraDaCompra := &ItemsCompra{
		Nome:       "Certamente não é manteiga",
		Quantidade: 1,
	}

	comprasDoMes := &ComprasMes{
		DataCompra:       time.Date(2023, 11, 23, 0, 0, 0, 0, time.Local),
		Mercado:          *localDeCompra,
		ItemsParaComprar: []ItemsCompra{*itemsCompraDaCompra},
	}

	return comprasDoMes
}
