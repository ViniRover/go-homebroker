# go-homebroker

## English

### Overview
**go-homebroker** is a trade order processing system built with Go and Kafka. It simulates a real-time order book, processing buy and sell orders efficiently.

### Features
- Uses **Kafka** for message streaming.
- Implements an **order book** to handle trade operations.
- Supports concurrent processing using **Go channels and goroutines**.

### Requirements
- [Go](https://go.dev/) installed.
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).
- Kafka running (can be set up via Docker).

### Installation
```sh
git clone https://github.com/ViniRover/go-homebroker.git
cd go-homebroker
docker-compose up -d
```

### Running the Application
```sh
go run cmd/trade/main.go
```

---

## Português

### Visão Geral
**go-homebroker** é um sistema de processamento de ordens de negociação desenvolvido em Go e Kafka. Ele simula um livro de ordens em tempo real, processando compras e vendas de forma eficiente.

### Funcionalidades
- Utiliza **Kafka** para streaming de mensagens.
- Implementa um **livro de ofertas** para operações de negociação.
- Suporta processamento concorrente com **canais e goroutines do Go**.

### Requisitos
- [Go](https://go.dev/) instalado.
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/).
- Kafka em execução (pode ser configurado via Docker).

### Instalação
```sh
git clone https://github.com/ViniRover/go-homebroker.git
cd go-homebroker
docker-compose up -d
```

### Executando a Aplicação
```sh
go run cmd/trade/main.go
```

