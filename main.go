package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Nome, Descricao string
	Quantidade, Id  int
	Preco           float64
}

// const (
// 	host     = "localhost"
// 	port     = "5433"
// 	user     = "postgres"
// 	password = "123123"
// 	dbName   = alura_loja
// )

func conectaComBancoDeDados() *sql.DB {
	connAluraLoja := ("host=localhost port=5433 user=postgres " + "password=123123 dbname=alura_loja sslmode=disable")
	db, err := sql.Open("postgres", connAluraLoja)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	lstProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for lstProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = lstProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Preco = preco
		p.Descricao = descricao
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4200", nil)
}
