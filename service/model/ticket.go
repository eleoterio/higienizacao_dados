package model

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Ticket local onde armazena os dados de importação do arquivo para salvar no banco:psql table:ticket
type Ticket struct {
	Cpf                string `json:"cpf" db:"cpf"`
	Private            string `json:"private" db:"private"`
	Incompleto         string `json:"incompleto" db:"incompleto"`
	DataUltimaCompra   string `json:"data_ultima_compra" db:"data_ultima_compra"`
	TicketMedio        string `json:"ticket_medio" db:"ticket_medio"`
	TicketUltimaCompra string `json:"ticket_ultima_compra" db:"ticket_ultima_compra"`
	LojaMaisFrequente  string `json:"loja_mais_frequente" db:"loja_mais_frequente"`
	LojaUltimaCompra   string `json:"loja_ultima_compra" db:"loja_ultima_compra"`
}

//GetCpf formatação string higienizada
func (i *Ticket) GetCpf() string {
	return cleanString(i.Cpf, 11)
}

//GetPrivate conversão do tipo 0 = false / 1= true
func (i *Ticket) GetPrivate() bool {
	if i.Private == "1" {
		return true
	}
	return false
}

//GetIncompleto conversão do tipo 0 = false / 1= true
func (i *Ticket) GetIncompleto() bool {
	if i.Incompleto == "1" {
		return true
	}
	return false
}

//SetTicketMedio formatação string to float original 2,66 resultado 2.66
func (i *Ticket) SetTicketMedio(ticketMedio string) {
	i.TicketMedio = strings.Replace(ticketMedio, ",", ".", -1)
}

//GetTicketMedio formatação string to float original 2,66 resultado 2.66
func (i *Ticket) GetTicketMedio() float64 {
	strOrigem := strings.Replace(i.TicketMedio, ",", ".", -1)
	floatResult, _ := strconv.ParseFloat(strOrigem, 64)
	return floatResult
}

//SetTicketUltimaCompra formatação string to float original 2,66 resultado 2.66
func (i *Ticket) SetTicketUltimaCompra(ticketUltimaCompra string) {
	i.TicketUltimaCompra = strings.Replace(ticketUltimaCompra, ",", ".", -1)
}

//GetTicketUltimaCompra formatação string to float original 2,66 resultado 2.66
func (i *Ticket) GetTicketUltimaCompra() float64 {
	strOrigem := strings.Replace(i.TicketUltimaCompra, ",", ".", -1)
	floatResult, _ := strconv.ParseFloat(strOrigem, 64)
	return floatResult
}

//GetDataUltimaCompra formatação da data
func (i *Ticket) GetDataUltimaCompra() (time.Time, bool) {
	var date time.Time
	if i.DataUltimaCompra == "NULL" {
		return date, false
	}
	date, err := time.Parse("2006-01-02", i.DataUltimaCompra)
	if err != nil {
		log.Println("[GetDataUltimaCompra] Erro ao converter string em Date: ", err)
		return date, false
	}
	return date, true
}

//GetLojaMaisFrequente formatação string higienizada
func (i *Ticket) GetLojaMaisFrequente() string {
	return cleanString(i.LojaMaisFrequente, 14)
}

//GetLojaUltimaCompra formatação string higienizada
func (i *Ticket) GetLojaUltimaCompra() string {
	return cleanString(i.LojaUltimaCompra, 14)
}

//GenerateTicket organiza dados vindo do arquivo para a inserção na table:"ticket"
func GenerateTicket(list []string) Ticket {
	var ticket Ticket
	ticket.Cpf = list[0]
	ticket.Private = list[1]
	ticket.Incompleto = list[2]
	ticket.DataUltimaCompra = list[3]
	ticket.SetTicketMedio(list[4])
	ticket.SetTicketUltimaCompra(list[5])
	ticket.LojaMaisFrequente = list[6]
	ticket.LojaUltimaCompra = list[7]

	return ticket
}

func cleanString(texto string, total int) string {
	regex := regexp.MustCompile("[0-9]+")
	texto = strings.Join(regex.FindAllString(texto, -1), "")
	for {
		texto = "0" + texto
		if len(texto) > total {
			return texto[0:total]
		}
	}
	return texto
}
