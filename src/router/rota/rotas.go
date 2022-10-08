package rota

import (
	"net/http"
	"github.com/gorilla/mux"


)

type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configura_Rota(r *mux.Router) *mux.Router{
	rotas := RotasUsuraios

	for _, rota := range rotas{
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	// fileserver Desbloqueia  css e javascript que contem na pasta /assets/
	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return r 
}