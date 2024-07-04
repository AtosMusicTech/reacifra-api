package _errors

type CifraNewError struct {
}

func NewCifraNewError() *CifraNewError {
	return &CifraNewError{}
}

func (e CifraNewError) Error() string {
	return "cifra erro"
}
