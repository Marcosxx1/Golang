package model

type Imovel struct {
	Tamanho  string
	Valor    float64
	NumeroDeQuartos int
	Endereco Endereco
}

type Casa struct {
	Imovel
	Jardim          bool
}

type Apartamento struct {
	Imovel
	Andar             int
	NumeroDeBlocos    int
	Elevador          bool
}
