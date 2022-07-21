package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBConnection = ""
	PORT         = 0
)

// Vai inicializar as variáveis de ambiente da aplicação:
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	PORT, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		PORT = 9000
	}

	// Seta string de conexão com o banco de dados:
	// O sprintf serve para formatar a string de conexão pegando os dados do .env
	DBConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

}
