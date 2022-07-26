package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	if erro := json.NewEncoder(res).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}

func Erro(res http.ResponseWriter, statusCode int, erro error) {
	JSON(res, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
