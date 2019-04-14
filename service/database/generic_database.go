package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	//Import para banco de dados
	_ "github.com/lib/pq"
)

//variavel singleton para armazenar a conexao
var db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL Abre a Conexão com o banco Postgress
func AbreConexaoComBancoDeDadosSQL() (db *sqlx.DB, err error) {
	err = nil
	db, err = sqlx.Open("postgres", "host=0.0.0.0 port=5432 user=postgres-dev dbname=neoway password=s3cr3tp4ssw0rd sslmode=disable")
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro na conexao: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro no ping na conexao: ", err.Error())
		return
	}
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
	err = db.Ping()
	if err != nil {
		log.Println("[GetDBConnection] Erro no ping na conexao: ", err.Error())
		return
	}
	localdb = db
	return
}

//CloseDBConnection Obtem a conexao com o banco de dados
func CloseDBConnection() (localdb *sqlx.DB, err error) {
	localdb.Close()
	return
}
