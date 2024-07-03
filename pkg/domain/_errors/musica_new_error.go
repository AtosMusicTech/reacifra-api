package _errors

type MusicaNewError struct {
}

func NewMusicaNewError() *MusicaNewError {
	return &MusicaNewError{}
}

func (e MusicaNewError) Error() string {
	return "musica nao encontrada"
}
