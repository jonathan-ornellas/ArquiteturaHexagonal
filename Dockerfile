# Usando a imagem oficial do Golang
FROM golang:1.16

# Definindo o diretório de trabalho
WORKDIR /go/src

# Habilitando CGO para utilizar bibliotecas C, como o sqlite3
ENV CGO_ENABLED=1
ENV PATH="/go/bin:${PATH}"

# Instalando dependências necessárias para desenvolvimento, incluindo gcc e sqlite3
RUN apt-get update && apt-get install -y \
    gcc \
    sqlite3 \
    libsqlite3-dev

# Instalando algumas ferramentas como cobra e mockgen
RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest && \
    export PATH=$PATH:$(go env GOPATH)/bin

# Alterando permissões do usuário e cache
RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache

# Definindo o usuário não-root para execução dos comandos
USER www-data

# Copiando o código para o container
COPY . .


CMD ["tail", "-f", "/dev/null"]
