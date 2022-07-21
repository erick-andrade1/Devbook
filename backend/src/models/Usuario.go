package models

import "time"

// O 'omitempty' irá fazer com que caso o campo ID esteja em branco quando for passar para JSON, o campo id não será passado no json
type user struct {
	ID        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
