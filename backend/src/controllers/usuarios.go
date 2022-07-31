package controllers

import (
	"backend/src/database"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	body, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		responses.Erro(res, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.User
	if erro = json.Unmarshal(body, &usuario); erro != nil {
		responses.Erro(res, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.PrepareUser(); erro != nil {
		responses.Erro(res, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(res, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	usuario.ID, erro = repository.Create(usuario)
	if erro != nil {
		responses.Erro(res, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(res, http.StatusCreated, usuario)
}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	nomeOuNick := strings.ToLower(req.URL.Query().Get("usuario"))

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(res, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, erro := repository.GetAllUsers(nomeOuNick)
	if erro != nil {
		responses.Erro(res, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(res, http.StatusFound, users)
}

func GetUserById(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Buscando usuário"))
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Atualizando usuário"))
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Deletando usuário"))
}
