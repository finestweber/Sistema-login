package controllers

import (
	"net/http"
	"go-webapp/src/utils"
	
)
func CarregaTelaLogin(w http.ResponseWriter, r *http.Request){
	//w.Write([]byte("Carregando tela de Login"))
	utils.ExecutarTemplate(w, "login.html", nil)
}