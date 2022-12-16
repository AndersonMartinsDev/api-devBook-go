package modelos

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()

	return nil
}

func (p *Publicacao) validar() error {
	if erro := p.validated(p.Titulo, "Titulo"); erro != nil {
		return erro
	}
	if erro := p.validated(p.Conteudo, "Conteudo"); erro != nil {
		return erro
	}
	return nil
}

func (p *Publicacao) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)
}

func (p *Publicacao) validated(text string, campo string) error {
	if text == "" {
		return errors.New(fmt.Sprintf("O %s não pode estar em branco", campo))
	}
	return nil
}
