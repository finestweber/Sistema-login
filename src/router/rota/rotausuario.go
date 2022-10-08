package rota

import (
	"net/http"

	"go-WebApp/src/controllers"
)

var RotasUsuraios = []Rota{
	{
		URI: "/",
		Metodo: http.MethodGet,
		Funcao: controllers.Carrega_Login,
		RequerAutenticacao: false,
	},
	{
		URI: "/login",
		Metodo: http.MethodGet,
		Funcao: controllers.Carrega_Login,
		RequerAutenticacao: false,
	},
}