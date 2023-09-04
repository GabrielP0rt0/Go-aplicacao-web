package services

import (
	"fmt"
	"module/db"
	"module/models"
)

func GetProdutos() []models.Produto {
	db := db.OpenDBAlura()

	lstProdutos, err := db.Query("select * from produtos order by id asc")
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

		p.Id = id
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

func Delete(id string) {
	db := db.OpenDBAlura()

	deleteDb, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic("#### Erro ao deletar dados no banco: " + err.Error())
	}

	deleteDb.Exec(id)
	defer db.Close()
}

func GetItem(id string) models.Produto {
	db := db.OpenDBAlura()

	query, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic("#### Erro ao selecionar produto para edição #### " + err.Error())
	}

	response := models.Produto{}
	for query.Next() {
		fmt.Println()
		err = query.Scan(&response.Id, &response.Nome, &response.Descricao, &response.Preco, &response.Quantidade)
		if err != nil {
			panic("#### Erro ao converter tabela para struct #### " + err.Error())
		}
		break
	}

	defer db.Close()
	return response
}

func UpdateProduto(p models.Produto) {
	db := db.OpenDBAlura()
	update, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic("#### Erro ao criar update no banco de dados #### " + err.Error())
	}

	update.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade, p.Id)
	defer db.Close()
}
