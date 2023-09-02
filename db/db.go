package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	postgresDriver = "postgres"
	host           = "localhost"
	port           = "5433"
	user           = "postgres"
	password       = "123123"
	dbName         = "alura_loja"
)

var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

func OpenDBAlura() *sql.DB {
	db, err := sql.Open(postgresDriver, dataSourceName)
	if err != nil {
		panic("#### Erro ao abrir conexão com o banco de dados: " + err.Error())
	}
	return db
}
