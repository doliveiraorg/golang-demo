# Use a imagem base do Go para compilar a aplicação
FROM golang:1.21.6-alpine AS builder

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copie os arquivos do projeto para o diretório de trabalho
COPY . .

# Baixe as dependências do módulo Go
RUN go mod download

# Compile a aplicação
RUN go build -o hello-world-api

# Use uma imagem base mínima para executar a aplicação
FROM alpine:latest

# Defina o diretório de trabalho dentro do container
WORKDIR /root/

# Copie o binário compilado do estágio anterior
COPY --from=builder /app/hello-world-api .

# Exponha a porta em que a aplicação será executada
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./hello-world-api"]

