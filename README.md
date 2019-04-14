# Neoway - Projeto para Higienização de dados


## Indice

* [Pré requisitos](#pré-requisitos)
* [Requisitos](#requisitos)
* [Desejável](#desejável)
* [Instalação](#instalação)
* [Configuração](#configuração)
* [Estrutura](#estrutura)
* [Autor](#autor)

## Pré requisitos
- Docker 
- GIT

## Requisitos:
- Criar um serviço em GO que receba um arquivo csv/txt de entrada (Arquivo Anexo)
- Este serviço deve persistir no banco de dados relacional (postgres) todos os dados contidos no arquivo
  Obs: O arquivo não possui um separador muito convencional
 
- Deve-se fazer o split dos dados em colunas no banco de dados
 Obs: pode ser feito diretamente no serviço em GO ou em sql
 
- Realizar higienização dos dados após persistência (sem acento, maiusculo, etc)
- Validar os CPFs/CNPJs contidos (validos e não validos numericamente)
- Todo o código deve estar disponível em repositório publico do GIT
 
## Desejável:
- Utilização das linguagen GOLANG para o desenvolvimento do serviço
- Utilização do DB Postgres
- Docker Compose , com orientações para executar (arquivo readme)

## Instalação
Após a instalação do docker e do git, baixar o projeto:
```
git clone https://github.com/eleoterio/neoway
```
entrar na pasta raiz do projeto e executar:
```
sudo docker-compose up
```
O projeto foi todo montado em cima do docker, por isso nao necessita nada instalado na maquina, basta baixar o projeto e subir o docker-compose

## Configuração
Acesso do banco:
```
POSTGRES_DB: dev
POSTGRES_USER: postgres-dev
POSTGRES_PASSWORD: s3cr3tp4ssw0rd
```
## Estrutura
Pastas:

- service (Projeto em go)
    - database (func relacionados a banco)
    - file (func relacionados ao manuseio do arquivo)
    - model (func relacionados ao modelos, tratamento de dados e validações)

Arquivos (dentro da pasta service):

- database (pasta)
    - databases_dev.sql (script de criação do banco de dados)
    - generic_database.go (func de abertura e fechamento do banco e criação das tabelas no banco, caso não exista)
    - inset_method.go (func para a insercção dos dados nas table "ticket" e "ticket_higienizado")
- file (pasta)
    - base_teste.txt (base de dados para teste do sistema)
    - read_file.go (abertura e leitura do arquivo para a inserção e tramento dos dados)
- model (pasta)
    - ticket_higienizado.go (criação da struct e geração dos dados para a inserção no banco)
    - ticket.go (criação da struct, validações dos campos para tratamento e higienização, geração dos dados para a inserção no banco)
    - validade (verificação de CPF, CNPJ e verificar string NULL, as func de validações de cpf e cnpj foi verificada na internet)
- main.go
    - Lê arquivo
    - Abre Conexão com o banco
    - Cria as tabelas caso não exista
    - Varre as linhas do arquivo
        - Insere na table "ticket" (original)
        - Insere na table "ticket_higienizado" (tratando as string, date e higienizando)
            - table "ticket_higienizada" os campos com "_valido" é a verificação se os dados informados estão validos ou não
    - Fechar a conexao com o banco

## Autor
* **[Felipe Eleoterio](https://www.linkedin.com/in/felipeeleoterio/)**