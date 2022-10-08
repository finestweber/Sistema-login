package controllers

import (
	"net/http"
	"go-WebApp/src/utils"
)

func Carrega_Login(w http.ResponseWriter, r *http.Request){
	utils.ExecutaTemplate(w, "login.html", nil)
}