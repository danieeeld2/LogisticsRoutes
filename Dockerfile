FROM golang:1.21-alpine3.17

LABEL mantainer="danieeeld2@correo.ugr.es"
LABEL version="5.0.0"

WORKDIR /app/test

COPY go.mod go.sum ./

RUN go mod tidy

RUN addgroup -S testgroup && adduser -S -u 1001 test -G testgroup

USER test

COPY internal internal

ENTRYPOINT ["go","test", "-v", "./internal"]
