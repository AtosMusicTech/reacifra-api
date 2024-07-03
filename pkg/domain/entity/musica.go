package entity

type Musica struct {
	Id     uint
	Titulo string
	Versao string
	Cifra  string
}

func NewMusica(id uint) *Musica {
	return &Musica{Id: id}
}
