# API em Go com PostgreSQL e Docker

A ideia do projeto é fornecer uma API com as funcionalidades mais básicas para realizar um sorteio.
Entre essas funcionalidades estão:
 * Criação usuários
 * Criação de sorteios
 * Geração de cupons
 * Sortear cupom

 ## 🚀 Começando

Essas instruções permitirão que você obtenha uma cópia do projeto em operação na sua máquina local para fins de desenvolvimento e teste.

### 🔧 Instalação

Para rodar o projeto em sua máquina, basta executar os seguintes comandos:

Subir os containers:

```
docker-compose up -d
```

Executar aplicação:

```
go run main.go
```


## 🛠️ Tecnologias utilizadas

* [Go lang](https://go.dev)
* [PostgreSQL](https://www.postgresql.org)
* [Docker](https://www.docker.com)