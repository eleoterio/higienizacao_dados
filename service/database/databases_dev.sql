CREATE TABLE ticket
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

CREATE TABLE ticket_higienizado
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
);
