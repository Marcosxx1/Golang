package model

type Imovel struct {
	Tamanho         string
	Valor           int
	NumeroDeQuartos int
	EnderecoImovel  Endereco
}

type Casa struct {
	Imovel
	Jardim bool
}

type Apartamento struct {
	Imovel
	Andar          int
	NumeroDeBlocos int
	Elevador       bool
}
