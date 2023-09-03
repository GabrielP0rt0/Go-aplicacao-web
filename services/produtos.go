package services

import (
	"module/db"
	"module/models"
)

func GetProdutos() []models.Produto {
	db := db.OpenDBAlura()

	lstProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := models.Produto{}
	produtos := []models.Produto{}

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

	defer db.Close()
	return produtos
}

func PostProduto(p models.Produto) {
	db := db.OpenDBAlura()

	insertDb, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic("#### Erro ao inserir dados no banco: " + err.Error())
	}

	insertDb.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)
	defer db.Close()
}
