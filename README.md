# Simple CRUD Go API

Uma API REST simples em **Golang** para gerenciamento de tarefas (tasks), desenvolvida por mim com o objetivo de **estudos** da linguagem Go e de conceitos de desenvolvimento de APIs.

## Sobre o projeto

Este projeto é um CRUD básico que permite listar e criar tarefas, persistindo os dados em um banco de dados **PostgreSQL**. Ele foi criado com foco no aprendizado de:

- Criação de APIs REST em Go
- Roteamento com [Gorilla Mux](https://github.com/gorilla/mux)
- Conexão e manipulação de banco de dados PostgreSQL com `database/sql`
- Variáveis de ambiente com [godotenv](https://github.com/joho/godotenv)

## Tecnologias utilizadas

- [Go](https://go.dev/) 1.26
- [Gorilla Mux](https://github.com/gorilla/mux) — roteador HTTP
- [lib/pq](https://github.com/lib/pq) — driver PostgreSQL
- [godotenv](https://github.com/joho/godotenv) — carregamento de variáveis de ambiente

## Estrutura do projeto

```
.
├── config/
│   └── db.go            # Configuração e conexão com o banco de dados
├── hadler/
│   └── task-handler.go  # Handlers das rotas de tasks
├── model/
│   └── task.go          # Modelo da entidade Task e SQL de criação da tabela
├── main.go              # Ponto de entrada da aplicação
└── .env.exemple         # Exemplo de configuração de variáveis de ambiente
```

## Endpoints

| Método | Rota     | Descrição                  |
|--------|----------|----------------------------|
| GET    | `/tasks` | Lista todas as tarefas     |
| POST   | `/tasks` | Cria uma nova tarefa       |

### Exemplo de requisição (POST /tasks)

```json
{
    "Title": "Estudar Go",
    "Description": "Aprender os fundamentos da linguagem",
    "status": false
}
```

### Exemplo de resposta (GET /tasks)

```json
[
    {
        "id": 1,
        "Title": "Estudar Go",
        "Description": "Aprender os fundamentos da linguagem",
        "status": false
    }
]
```

## Como rodar o projeto

### Pré-requisitos

- [Go](https://go.dev/dl/) instalado
- [PostgreSQL](https://www.postgresql.org/) em execução

### Passos

1. Clone o repositório:

```bash
git clone https://github.com/alnszzx/simple-crud--go.git
cd simple-crud--go
```

2. Configure as variáveis de ambiente. Renomeie o arquivo `.env.exemple` para `.env` e informe sua string de conexão com o PostgreSQL:

```env
DB_CONNECTION=sua url de conexão
```

3. Instale as dependências:

```bash
go mod tidy
```

4. Execute a aplicação:

```bash
go run main.go
```

A API estará disponível em `http://localhost:8080`.

> A tabela `tasks` é criada automaticamente na inicialização da aplicação, caso não exista.

---

Projeto desenvolvido para fins de estudo. 🚀
