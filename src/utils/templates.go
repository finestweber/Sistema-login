package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func CarregaTemplate(){
	templates = template.Must(template.ParseGlob("views/*.html"))
}
func ExecutaTemplate(w http.ResponseWriter, template string, dados interface{}){
	 templates.ExecuteTemplate(w, template, dados)
}