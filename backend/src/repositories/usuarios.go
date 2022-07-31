package repositories

import (
	"backend/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

// Puxa o banco e coloca no repositorio para que não seja preciso ficar chamando o banco sempre a cada metodo
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO users (nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Nome, user.Nick, user.Email, user.Senha)
	if erro != nil {
		return 0, erro
	}

	lastInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertedId), nil
}

func (repository users) GetAllUsers(nomeOuNick string) ([]models.User, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // fica algo tipo -> %nomeOuNick%

	lines, erro := repository.db.Query("SELECT id, nome, nick, email, createdAt FROM users WHERE nome LIKE ? or nick LIKE ?", nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		//Escaneia a linha e repassa para o usuario criado
		if erro = lines.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		//Insere o usuario criado no array de usuarios q ira retornar na chamada.
		users = append(users, user)
	}

	return users, nil
}
