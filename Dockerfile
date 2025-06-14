# Use a imagem oficial do Go como base
FROM golang:1.23-alpine AS builder

# Diretório de trabalho dentro do container
WORKDIR /app

# Copiar arquivos de modulação para cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o código da aplicação
COPY ./api ./api

# Build da aplicação
RUN go build -o goapi ./api

# Imagem final, menor e mais enxuta
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário da fase builder
COPY --from=builder /app/goapi .

# Expõe a porta da aplicação
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./goapi"]
