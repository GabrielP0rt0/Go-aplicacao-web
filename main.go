package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4200", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{{"Camiseta", "Azul, bem bonita", 39., 1},
		{"Tenis", "Confortável", 89., 3},
		{"Fone", "Para orelha", 100.99, 1}}
	temp.ExecuteTemplate(w, "Index", produtos)
}
