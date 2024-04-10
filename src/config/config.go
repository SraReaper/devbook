package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//AAPIURL representa a URL para a comunicação com a API
	APIURL   = ""
	Porta    = 0
	HashKey  []byte
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	//populando a variável porta
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	//populando a variável APIURL
	APIURL = os.Getenv("API_URL")

	//populando a variável HashKey
	HashKey = []byte(os.Getenv("HASH_KEY"))

	//populando a variável BlockKey
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
