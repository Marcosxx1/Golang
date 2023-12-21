package model

type Retangulo struct {
	Largura, Altura float64
}

func IniciaRetangulo() *Retangulo {

	retangulo := &Retangulo{
		Largura: 1,
		Altura:  2,
	}
	return retangulo
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Largura
}
