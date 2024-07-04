package _errors

type CifraUpdateError struct {
}

func NewCifraUpdateError() *CifraUpdateError {
	return &CifraUpdateError{}
}

func (e CifraUpdateError) Error() string {
	return "Erro ao atualizar a cifra"
}
