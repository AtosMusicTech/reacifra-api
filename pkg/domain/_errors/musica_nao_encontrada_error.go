package _errors

type MusicaNaoEncontradaError struct {
}

func NewMusicaNaoEncontradaError() *MusicaNaoEncontradaError {
	return &MusicaNaoEncontradaError{}
}

func (e MusicaNaoEncontradaError) Error() string {
	return "musica nao encontrada"
}
