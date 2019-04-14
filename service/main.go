package main

import (
	"encoding/json"
	"log"

	"github.com/eleoterio/neoway/service/database"
	"github.com/eleoterio/neoway/service/file"
)

//FILE caminho/arquivo csv ou txt a ser lido
const (
	FILE = "service/file/base_teste.txt"
)

func main() {
	list, err := file.ReadFile(FILE)
	if err != nil {
		log.Println("[main] Error ao ler o arquivo: ", err.Error())
		return
	}
	database.GetDBConnection()
	for _, lineColumn := range list {
		validade, ticket := database.InsertTicket(lineColumn)
		if validade {
			database.InsertTicketHigienizado(ticket)
		} else {
			json, _ := json.Marshal(ticket)
			log.Println("[main] Error ao inserir ticket: ", json)
		}
	}
	database.CloseDBConnection()
}
