package controllers

import (
	"html/template"
	"module/services"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := services.GetProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}
