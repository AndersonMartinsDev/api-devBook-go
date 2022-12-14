package modelos

import (
	"api/src/seguranca"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (u *Usuario) validar(etapa string) error {

	if erro := validated(u.Nome, "nome", etapa); erro != nil {
		return erro
	}

	if erro := validated(u.Nick, "nick", etapa); erro != nil {
		return erro
	}

	if erro := validated(u.Email, "email", etapa); erro != nil {
		return erro
	}

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("Email não é válido")
	}

	if erro := validated(u.Senha, "senha", etapa); erro != nil {
		return erro
	}
	return nil
}

func validated(text string, campo string, etapa string) error {
	if etapa == "cadastro" && text == "" {
		return errors.New(fmt.Sprintf("O %s não pode estar em branco", campo))
	}
	return nil
}

func (u *Usuario) formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(u.Senha)
		if erro != nil {
			return erro
		}

		u.Senha = string(senhaComHash)
	}

	return nil
}
