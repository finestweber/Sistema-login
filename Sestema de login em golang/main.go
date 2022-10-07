package main

import (
	"fmt"
	"go-webapp/src/utils"
	"go-webapp/src/config"
	"go-webapp/src/router"
	"log"
	"net/http"
)
func main(){
	config.Iniciar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Servidor Rodando porta :3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}