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
			log.Println("#### Erro na conversão do preço")
		}

		quantidade, err := strconv.Atoi(sQuantidade)
		if err != nil {
			log.Println("#### Erro na conversão da quantidade")
		}
		request := models.Produto{Nome: nome, Descricao: descricao, Quantidade: quantidade, Preco: preco}
		services.PostProduto(request)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	services.Delete(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	produto := services.GetItem(r.URL.Query().Get("id"))
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sId := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		sPreco := r.FormValue("preco")
		sQuantidade := r.FormValue("quantidade")

		preco, err := strconv.ParseFloat(sPreco, 64)
		if err != nil {
			log.Println("#### Erro na conversão do preço #### " + err.Error())
		}

		quantidade, err := strconv.Atoi(sQuantidade)
		if err != nil {
			log.Println("#### Erro na conversão da quantidade #### " + err.Error())
		}

		id, err := strconv.Atoi(sId)
		if err != nil {
			log.Println("#### Erro na conversão do id #### " + err.Error())
		}
		request := models.Produto{nome, descricao, quantidade, id, preco}
		services.UpdateProduto(request)
	}
	http.Redirect(w, r, "/", 301)
}
