package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	//Import para banco de dados
	_ "github.com/lib/pq"
)

//variavel singleton para armazenar a conexao
var db *sqlx.DB

//scrit de criação do banco, igual do databases_dev.sql
var schema = `
CREATE TABLE IF NOT EXISTS ticket
(
	cpf 					text NOT NULL,
	private 				text,
	incompleto 				text,
	data_ultima_compra 		text,
	ticket_medio 			text,
	ticket_ultima_compra 	text,
	loja_mais_frequente 	text,
	loja_ultima_compra 		text 
);

CREATE TABLE IF NOT EXISTS ticket_higienizado
(
	cpf 						text NOT NULL,
	private 					bool,
	incompleto 					bool,
	data_ultima_compra 			date,
	ticket_medio 				float,
	ticket_ultima_compra 		float,
	loja_mais_frequente 		text,
	loja_ultima_compra 			text,
	cpf_valido 					bool,
	loja_mais_frequente_valido 	bool,
	loja_ultima_compra_valido 	bool
);`

//AbreConexaoComBancoDeDadosSQL Abre a Conexão com o banco Postgress
func AbreConexaoComBancoDeDadosSQL() (db *sqlx.DB, err error) {
	err = nil
	db, err = sqlx.Connect("postgres", "host=db port=5432 user=postgres-dev dbname=dev password=s3cr3tp4ssw0rd sslmode=disable")
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro na conexao: ", err.Error())
		return
	}
	//Cria as tabelas no banco caso não exista
	db.MustExec(schema)
	return
}

//GetDBConnection Obtem a conexao com o banco de dados
func GetDBConnection() (localdb *sqlx.DB, err error) {
	if db == nil {
		db, err = AbreConexaoComBancoDeDadosSQL()
		if err != nil {
			log.Println("[GetDBConnection] Erro na conexao: ", err.Error())
			return
		}
	}
	localdb = db
	return
}

//CloseDBConnection Obtem a conexao com o banco de dados
func CloseDBConnection(localdb *sqlx.DB) {
	localdb.Close()
}
