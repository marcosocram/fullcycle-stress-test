# Use uma imagem base do Go
FROM golang:alpine as builder

# Defina o diretório de trabalho
WORKDIR /app

# Copie o código fonte para o container
COPY . .

# Instale as dependências
RUN go mod tidy

# Compile o executável
RUN go build -o fullcycle-stress-test .

# Defina o ponto de entrada
ENTRYPOINT ["/app/fullcycle-stress-test"]