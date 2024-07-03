package responses

import "cifra-api/pkg/domain/entity"

type MusicaResponse struct {
	Id     uint   `json:"id"`
	Titulo string `json:"titulo"`
	Versao string `json:"versao"`
	Cifra  string `json:"cifra"`
}

func ConvertToMusicasResponse(musica []*entity.Musica) []*MusicaResponse {
	musicasResponse := []*MusicaResponse{}

	for _, m := range musica {
		musicasResponse = append(musicasResponse, ConvertToMusicaResponse(m))
	}

	return musicasResponse
}

func ConvertToMusicaResponse(musica *entity.Musica) *MusicaResponse {
	return &MusicaResponse{
		Id:     musica.Id,
		Titulo: musica.Titulo,
		Versao: musica.Versao,
		Cifra:  musica.Cifra,
	}
}
