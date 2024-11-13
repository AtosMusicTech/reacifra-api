package entity

type Queue struct {
	Id     uint
	Cifras []*Cifra
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) AddCifra(cifra *Cifra) {
	q.Cifras = append(q.Cifras, cifra)
}
