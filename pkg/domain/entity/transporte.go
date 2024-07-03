package entity

type Transporte struct {
	Id      uint
	Musica  *Musica
	Posicao uint
}

func NewTransporte(id uint) *Transporte {
	return &Transporte{
		Id: id,
	}
}
