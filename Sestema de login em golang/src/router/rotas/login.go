package rotas

import (
	"net/http"
	"go-webapp/src/controllers"
	)
var RotasLogin = []Rota{
	{
	URI: "/",
	Metodo: http.MethodGet,
	Funcao: controllers.CarregaTelaLogin,
	RequerAutenticacao: false,
	},
	{
	URI: "/login",
	Metodo: http.MethodGet,
	Funcao: controllers.CarregaTelaLogin,
	RequerAutenticacao: false,
	},
}
