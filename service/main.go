package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/eleoterio/neoway/service/database"
	"github.com/eleoterio/neoway/service/file"
)

//FILE caminho/arquivo csv ou txt a ser lido
const (
	FILE = "service/file/base_teste.txt"
)

func main() {
	/*
		O banco demora mais que o GO para iniciar, por isso este sleep;
		Caso você inicie primeiro o postgres e depois rode o programa, pode ser retirada essa linha
	*/
	time.Sleep(2 * time.Second)
	list, err := file.ReadFile(FILE)
	if err != nil {
		log.Println("[main] Error ao ler o arquivo: ", err.Error())
		return
	}
	localdb, err := database.GetDBConnection()
	if err != nil {
		log.Println("[main] Error na conexao com o banco: ", err.Error())
		return
	}
	defer database.CloseDBConnection(localdb)
	fmt.Println("Inicio do Processo de importação e persistencia de dados no banco!")
	for _, lineColumn := range list {
		validade, ticket := database.InsertTicket(lineColumn)
		if validade {
			database.InsertTicketHigienizado(ticket)
		} else {
			json, _ := json.Marshal(ticket)
			log.Println("[main] Error ao inserir ticket: ", json)
		}
	}
	fmt.Println("Fim do Processo de importação e persistencia de dados no banco!")

}
