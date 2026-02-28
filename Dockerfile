# Build
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Aqui você gera o executável chamado 'meu-app' na raiz /app
RUN CGO_ENABLED=0 GOOS=linux go build -o meu-app ./cmd/main.go

# Execution
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# CORREÇÃO AQUI: Copiamos o arquivo 'meu-app' da raiz /app do builder
COPY --from=builder /app/meu-app .

# Copia o config (garanta que o arquivo existe na sua pasta local)
COPY config-file.ini .

EXPOSE 8080

# O comando deve apontar exatamente para o nome do arquivo que copiamos acima
CMD ["./meu-app"]
