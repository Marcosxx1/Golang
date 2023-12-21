package model

import "math"

type Circulo struct {
	Raio float64
}

func IniciaCirculo() *Circulo {

	circulo := &Circulo{
		Raio: 3,
	}
	return circulo
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
