package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexaoBanco é a string de conexão com o MySql
	StringConexaoBanco = ""

	//Porta onde a API vai estar rodando
	Porta = 0

	//SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 9000
	}

	db_usuario := os.Getenv("DB_USUARIO")
	db_senha := os.Getenv("DB_SENHA")
	db_nome := os.Getenv("DB_NOME")
	db_config := "charset=utf8&parseTime=True&loc=Local"
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?%s", db_usuario, db_senha, db_nome, db_config)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
