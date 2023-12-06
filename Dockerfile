FROM golang:1.21-alpine3.17

LABEL mantainer="danieeeld2@correo.ugr.es" \
      version="5.0.0"

WORKDIR /app

RUN adduser -D -u 1001 test && \
    chown test /app && \
    apk add build-base

USER test

COPY go.mod go.sum magefile.go ./

RUN go mod download && \
    go install github.com/magefile/mage@v1.15.0 && \
    mage -compile mageCompilado && \
    rm go.sum magefile.go go.mod

WORKDIR /app/test

ENTRYPOINT [ "../mageCompilado", "test"]
