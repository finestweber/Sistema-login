package main
 
import (
    "net/http"
    "fmt"
    "go-WebApp/config"
    "log"
    "go-WebApp/src/router"
    "go-WebApp/src/utils"
)
func main(){
    config.Iniciar()
    utils.CarregaTemplate()

    fmt.Println("Servidor Rodando Porta : 5000")

    r := router.Carrega_Router()
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
 
 }