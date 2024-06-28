# Etapa de build
FROM golang:1.22.3 AS builder

# Configura o diretório de trabalho dentro do container
WORKDIR /app

# Copia o go.mod e go.sum e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ecommerce-carrinho .

# Etapa de produção
FROM scratch

# Copia o binário compilado da etapa de build
COPY --from=builder /app/ecommerce-carrinho /ecommerce-carrinho

# Define a porta que o aplicativo vai expor
EXPOSE 8082

# Define o ponto de entrada do container
ENTRYPOINT ["/ecommerce-carrinho"]