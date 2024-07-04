package _errors

type CifraNaoEncontradaError struct {
}

func NewCifraNaoEncontradaError() *CifraNaoEncontradaError {
	return &CifraNaoEncontradaError{}
}

func (e CifraNaoEncontradaError) Error() string {
	return "cifra não encontrada"
}
