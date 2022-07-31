package models

import (
	"errors"
	"strings"
	"time"
)

// O 'omitempty' irá fazer com que caso o campo ID esteja em branco quando for passar para JSON, o campo id não será passado no json
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Irá validar e formatar o usuário antes de inserir no banco de dados:
func (user *User) PrepareUser() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco!")
	}
	if user.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco!")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco!")
	}
	if user.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco!")
	}

	return nil
}

func (user *User) format() {
	user.Nome = strings.TrimSpace(user.Nome)
	user.Nick = strings.TrimSpace(user.Nome)
	user.Email = strings.TrimSpace(user.Nome)
}
