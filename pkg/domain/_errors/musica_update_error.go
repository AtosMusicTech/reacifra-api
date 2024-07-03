package _errors

type MusicaUpdateError struct {
}

func NewMusicaUpdateError() *MusicaUpdateError {
	return &MusicaUpdateError{}
}

func (e MusicaUpdateError) Error() string {
	return "Erro ao atualizar a música"
}
