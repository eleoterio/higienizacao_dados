package model

//Banco de dados Postgress
import (
	"database/sql"

	"github.com/lib/pq"
)

//TicketHigienizado local onde armazena os dados ja higienizado banco:psql
type TicketHigienizado struct {
	Cpf                     string          `json:"cpf" db:"cpf"`
	Private                 bool            `json:"private" db:"private"`
	Incompleto              bool            `json:"incompleto" db:"incompleto"`
	DataUltimaCompra        pq.NullTime     `json:"data_ultima_compra" db:"data_ultima_compra"`
	TicketMedio             sql.NullFloat64 `json:"ticket_medio" db:"ticket_medio"`
	TicketUltimaCompra      sql.NullFloat64 `json:"ticket_ultima_compra" db:"ticket_ultima_compra"`
	LojaMaisFrequente       sql.NullString  `json:"loja_mais_frequente" db:"loja_mais_frequente"`
	LojaUltimaCompra        sql.NullString  `json:"loja_ultima_compra" db:"loja_ultima_comprapf"`
	CpfValido               bool            `json:"cpf_valido" db:"cpf_valido"`
	LojaMaisFrequenteValido bool            `json:"loja_mais_frequente_valido" db:"loja_mais_frequente_valido"`
	LojaUltimaCompraValido  bool            `json:"loja_ultima_compra_valido" db:"loja_ultima_comprapf_valido"`
}

//GenerateTicketHigienizado organiza dados vindo do arquivo para a inserção na table:"ticket"
func GenerateTicketHigienizado(ticket Ticket) TicketHigienizado {
	var ticketHigienizado TicketHigienizado
	ticketHigienizado.Cpf = ticket.GetCpf()
	ticketHigienizado.Private = ticket.GetPrivate()
	ticketHigienizado.Incompleto = ticket.GetIncompleto()
	if date, valida := ticket.GetDataUltimaCompra(); valida {
		ticketHigienizado.DataUltimaCompra = pq.NullTime{date, true}
	}
	if ValidaNull(ticket.TicketMedio) {
		ticketHigienizado.TicketMedio = sql.NullFloat64{ticket.GetTicketMedio(), true}
	}
	if ValidaNull(ticket.TicketUltimaCompra) {
		ticketHigienizado.TicketUltimaCompra = sql.NullFloat64{ticket.GetTicketUltimaCompra(), true}
	}
	if ValidaNull(ticket.LojaMaisFrequente) {
		ticketHigienizado.LojaMaisFrequente = sql.NullString{ticket.GetLojaMaisFrequente(), true}
	}
	if ValidaNull(ticket.LojaUltimaCompra) {
		ticketHigienizado.LojaUltimaCompra = sql.NullString{ticket.GetLojaUltimaCompra(), true}
	}
	ticketHigienizado.CpfValido = ValidaCPF(ticket.GetCpf())
	ticketHigienizado.LojaMaisFrequenteValido = ValidaCNPJ(cleanString(ticket.LojaMaisFrequente, 14))
	ticketHigienizado.LojaUltimaCompraValido = ValidaCNPJ(cleanString(ticket.LojaUltimaCompra, 14))

	return ticketHigienizado
}
