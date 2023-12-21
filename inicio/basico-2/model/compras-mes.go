/*
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
			Mercado:          *localDeCompra,                      //Aponta para o valor de Mercado criado localmente
			ItemsParaComprar: []ItemsCompra{*itemsCompraDaCompra}, //Aponta para o valor de ItemsCompr criado localmente
		}

		return comprasDoMes
	}
*/
package model

import "time"

type ComprasMes struct {
	DataCompra       time.Time
	Mercado          *Mercado
	ItemsParaComprar []*ItemsCompra
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
		Mercado:          localDeCompra,                       //Aponta para uma instancia existente de Mercado
		ItemsParaComprar: []*ItemsCompra{itemsCompraDaCompra}, // Aponta oara uma instancia existente de ItemsCompra
	}

	return comprasDoMes
}
