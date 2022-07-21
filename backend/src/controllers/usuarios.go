package controllers

import "net/http"

func CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Criando usuário"))
}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Buscando todos os usuários"))
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
