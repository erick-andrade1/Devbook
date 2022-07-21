package router

import (
	"backend/src/router/rotas"

	"github.com/gorilla/mux"
)

// Função responsável por gerar as rotas e retorná-las
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
