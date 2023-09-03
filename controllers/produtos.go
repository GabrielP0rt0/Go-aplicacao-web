package controllers

import (
	"html/template"
	"log"
	"module/models"
	"module/services"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := services.GetProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		sPreco := r.FormValue("preco")
		sQuantidade := r.FormValue("quantidade")

		preco, err := strconv.ParseFloat(sPreco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
		}

		quantidade, err := strconv.Atoi(sQuantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade")
		}
		request := models.Produto{Nome: nome, Descricao: descricao, Quantidade: quantidade, Preco: preco}
		services.PostProduto(request)
	}
	http.Redirect(w, r, "/", 301)
}
