# goapi

API simples em Go voltada para fins educacionais e experimentação com práticas modernas de DevOps, incluindo:

- Build e versionamento de imagem Docker
- Deploy via Kubernetes
- Gestão de infraestrutura com Terraform
- Automatização com GitHub Actions

## 🎯 Propósito

Este projeto tem como objetivo servir de base para aprendizado e demonstração de boas práticas de:

- Organização de projetos em Go
- Empacotamento com Docker
- Deploy em ambientes orquestrados (como Kubernetes)
- Separação entre código de aplicação e infraestrutura

A aplicação em si é uma API CRUD básica, que se conecta a um banco PostgreSQL e permite gerenciar registros simples.

## 📦 Estrutura

```

goapi/
├── Dockerfile             # Define a imagem da API
├── go.mod / go.sum        # Dependências Go
├── api/                   # Ponto de entrada da aplicação
│   └── main.go
├── db/                    # Conexão com banco de dados
├── handlers/              # Lógica de tratamento de rotas
├── models/                # Estrutura dos dados
└── router/                # Definição das rotas da API

````

## 🚀 Executar localmente

### Com Docker Compose

```bash
docker compose up --build
````

API disponível em `http://localhost:8080`.

### Sem Docker

```bash
go run ./api
```

(Requer banco de dados PostgreSQL rodando localmente)

## 🐳 Docker

### Build manual da imagem

```bash
docker build -t ghcr.io/seu-usuario/goapi:<tag> .
```

### Publicação (feita via CI)

Ao criar uma nova tag Git (`v1.2.3`), o CI automaticamente:

* Faz o build da imagem
* Publica para: `ghcr.io/seu-usuario/goapi:<tag>`
* Usa tags imutáveis para rastreabilidade e rollback seguro

## 🧪 Testes

```bash
go test ./...
```

## 📘 RFCs

Todas as decisões de arquitetura, organização e deploy estão documentadas na pasta [`docs/rfcs/`](./docs/rfcs/).

Documento inicial: [`docs/rfcs/0001-split-between-API-and-infra.md.md`](./docs/rfcs/0001-split-between-API-and-infra.md.md)

## 🗂️ Repositórios relacionados

Este repositório contém **apenas o código da API**.

Infraestrutura (Kubernetes, Terraform, banco etc.) está no repositório:

👉 [`goapi-infra`](https://github.com/seu-usuario/goapi-infra)
