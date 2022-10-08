package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Porta = 0

func Iniciar(){
	var erro error

	if erro = godotenv.Load(); erro != nil{
		log.Fatal(erro)
	}

	if Porta, erro = strconv.Atoi(os.Getenv("WEB_APP_PORT")); erro != nil{
		Porta = 1234
	}
}