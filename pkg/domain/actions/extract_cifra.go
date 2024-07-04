package actions

import "cifra-api/pkg/domain"

type ExtractCifraAction struct {
	request domain.Requests
}

func NewExtractCifraAction() *ExtractCifraAction {
	return &ExtractCifraAction{}
}

func (e *ExtractCifraAction) Execute(url string) (string, error) {
	return e.request.Get(url)
}
