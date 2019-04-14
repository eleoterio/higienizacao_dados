package database

import (
	"log"

	"github.com/eleoterio/neoway/service/model"
)

//InsertTicket inseri o ticket no banco de dados
func InsertTicket(list []string) (bool, model.Ticket) {
	ticket := model.GenerateTicket(list)

	sqlInsert := "INSERT INTO ticket (" +
		"cpf," +
		"private," +
		"incompleto," +
		"data_ultima_compra," +
		"ticket_medio," +
		"ticket_ultima_compra," +
		"loja_mais_frequente," +
		"loja_ultima_compra" +
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
	_, err := db.Exec(sqlInsert,
		ticket.Cpf,
		ticket.Private,
		ticket.Incompleto,
		ticket.DataUltimaCompra,
		ticket.TicketMedio,
		ticket.TicketUltimaCompra,
		ticket.LojaMaisFrequente,
		ticket.LojaUltimaCompra)
	if err != nil {
		log.Println("[InsertTicket] Erro na insercao na tabela ticket: ", err)
		return false, ticket
	}
	return true, ticket
}

//InsertTicketHigienizado inseri o ticket ja tratado no banco de dados
func InsertTicketHigienizado(ticket model.Ticket) bool {
	ticketHigienizado := model.GenerateTicketHigienizado(ticket)

	sqlInsert := "INSERT INTO ticket_higienizado (" +
		"cpf," +
		"private," +
		"incompleto," +
		"data_ultima_compra," +
		"ticket_medio," +
		"ticket_ultima_compra," +
		"loja_mais_frequente," +
		"loja_ultima_compra," +
		"cpf_valido," +
		"loja_mais_frequente_valido," +
		"loja_ultima_compra_valido" +
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);"
	_, err := db.Exec(sqlInsert,
		ticketHigienizado.Cpf,
		ticketHigienizado.Private,
		ticketHigienizado.Incompleto,
		ticketHigienizado.DataUltimaCompra,
		ticketHigienizado.TicketMedio,
		ticketHigienizado.TicketUltimaCompra,
		ticketHigienizado.LojaMaisFrequente,
		ticketHigienizado.LojaUltimaCompra,
		ticketHigienizado.CpfValido,
		ticketHigienizado.LojaMaisFrequenteValido,
		ticketHigienizado.LojaUltimaCompraValido)
	if err != nil {
		log.Println("[InsertTicketHigienizado] Erro na insercao na tabela ticket: ", err)
		return false
	}
	return true
}
