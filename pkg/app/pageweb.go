package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageWeb struct {
	doc *goquery.Document
}

func NewPageWeb() *PageWeb {
	return &PageWeb{}
}

// fetchHTML faz uma requisição HTTP para obter o HTML da página na URL fornecida
func (r *PageWeb) Load(url string) error {
	// Faz a requisição HTTP
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erro ao fazer a requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verifica se a resposta HTTP foi bem-sucedida
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status HTTP não OK: %s", resp.Status)
	}

	// Carrega o HTML na estrutura do goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao carregar o HTML: %v", err)
	}

	r.doc = doc

	return nil
}

// extractPreContent extrai o conteúdo de todos os elementos <pre> do documento HTML
func (r *PageWeb) ExtractHtmlByTagName(tagname string) string {
	var contents []string

	// Seleciona todos os elementos <pre> e extrai o texto de cada um
	r.doc.Find(tagname).Each(func(i int, s *goquery.Selection) {
		content, _ := s.Html()
		contents = append(contents, content)
	})

	return strings.Join(contents, "")
}
