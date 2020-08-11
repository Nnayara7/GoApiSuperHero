# Documentação super-hero-api
  Projeto Api Levpay
# Organização dos Diretórios:
Pasta Principal é go-api/ , lá contém o arquivo (api.go) que é o arquivo para executar as chamadas de metódos já com o teste de conexão. 

-> Diretório go-api/Models 

  Dentro do diretório contém o arquivos connectio.go  que vai conectar a API(superheroapi.com) com o Banco de Dados (Postresql)


-> Diretório go-api/Models 

  contém o arquivo super.go que tem funções para manipular o banco de dados

-> Diretório Super-Hero-Api/types
 
  Contém o arquivo types.go que são os tipos usados ex:(types.Get usado para receber os dados do Banco)

# Funcionamento dos Arquivos
# Caminho : /routes/routes.go

O pakcage routes tem acesso as funções AddSuper, ListSuper, FindSuperByName, FindSuperById, RemoveSuper  do controller SuperController. 

O código é compativel para testagem via Postman


# Caminho: models/connections.go 

imports de fora do projeto:

"database/sql" -> para conectar com o banco

"io/ioutill" -> usado para ler o json da API(superheroapi.com) 

"fmt" -> padrão para input/output

"net/http" -> para fazer requisições do http da API(superheroapi.com)

AddSuper()-> Nessa função vc cria um cadastro, não exige paramêtro mas necessita de um nome para busca na API

ListSuper() -> Busca pelo tipo do super (good, bad)

FindSuperByName() -> Busca pelo nome do super

FindSuperById() Busca pelo id do super

RemoveSuper()-> Recebe uma String "name", e excluí o super indicado


# Caminho: models/super.go
importações de fora do projeto 
"database/sql" -> para uso no postresql.
"github.com/lib/pq" -> utilizando para colocar array do tipo json

Funções 

Saving()-> Recebe dois parâmetros uma slice de (id's) que foram passsados no arquivo (connection.go) esses id's são resultados obtidos na consulta http com o nome digitados pelo usuário.


checkValidityId() -> essa função irá avaliar  pelo id se o super correspondente poderá ser cadrastado. Retorna dois valores, true para sinalizar que algum super pode ser cadrastado

NewSuper() -> recebe três parâmetros (db *sql.DB, super *types.AllSuper, posResponse int) db para acesso as funções do banco , super que é a estrutura que contém todos os dados do SUPER a ser cadrastado e posResponse que a posição obtida pela função checkValidityId(). A função simplemente irá passar os dados da estrutura para as tabelas existentes.

GetSuperById() -> pesquisa um super por seu ID

GetSuperName() -> pesquisa um super por seu nome

GetSuperByAlignment() -> pesquisa super por caracteristica (bad, good(vilão ou herói ))

RemoveDataBase() -> faz um drop da tabela
 
DeleteSuper() -> deleta o super por seu nome 

# Executando  

Requistos : Necessário criar um server/banco de dados 

rodar todos os create table do arquivo sql.sql

Vá até a pasta go-api e rode o seguinte comando

go run api.go -config ./configs/config.yml

Assim ele irá rodar a linha descomentada do arquivo routes/routes.go

# Testes

para rodar os testes bastar ir até a basta tests e rodar o seguinte comando
go test ./... 

Também pode ser feito testes atráves da ferramenta postman:

http://localhost:3000/super
POST ->
	Body:
		{
    		"super-name": "Batman"
		}
http://localhost:3000/superName
GET ->
	Body:
		{
    		"id": "70"
		}

http://localhost:3000/superId
GET ->
	Body:
		{
    		"super-name": "Batman"
		}

http://localhost:3000/super
GET ->
	Body:
		{
    		"type": "good"
		}

http://localhost:3000/super
DELETE ->
	Body:
		{
    		"super-name": "Batman"
		}

# Especificações
go version go1.14.2 linux/amd64

Postresql 10.12

Sistema Operacional: linux Ubuntu 18.4

# Importações
go get github.com/steinfletcher/apitest
go get github.com/steinfletcher/apitest-jsonpath
go get -u github.com/gorilla/mux

# Contatos
nnayara.pedrozo@gmail.com
