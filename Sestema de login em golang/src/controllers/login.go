package controllers

import (
	"net/http"
	"go-webapp/src/utils"
	
)
func CarregaTelaLogin(w http.ResponseWriter, r *http.Request){
	utils.ExecutarTemplate(w, "login.html", nil)
}