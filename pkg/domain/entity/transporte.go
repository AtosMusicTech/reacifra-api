package entity

type Transporte struct {
	Id      uint
	Cifra   *Cifra
	Posicao uint
}

func NewTransporte(id uint) *Transporte {
	return &Transporte{
		Id: id,
	}
}
