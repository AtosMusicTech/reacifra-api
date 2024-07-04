package entity

type Cifra struct {
	Id         uint
	Titulo     string
	Versao     string
	Texto      string
	Tonalidade string
}

func NewCifra(id uint) *Cifra {
	return &Cifra{Id: id}
}
