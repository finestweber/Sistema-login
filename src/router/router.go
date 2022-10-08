package router

import (
	"go-WebApp/src/router/rota"
	"github.com/gorilla/mux"
)
func Carrega_Router() *mux.Router{
	r := mux.NewRouter()
	return rota.Configura_Rota(r)
}
