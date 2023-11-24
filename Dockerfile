FROM golang:1.21-alpine3.17

LABEL mantainer="danieeeld2@correo.ugr.es"
LABEL version="5.0.0"

RUN addgroup -S testgroup && adduser -S -u 1001 test -G testgroup
WORKDIR /user/test

RUN apk update && apk add git
RUN go install github.com/magefile/mage@v1.15.0
RUN chown -R test:testgroup /user/test

USER test

COPY go.mod .
COPY go.sum .
COPY magefile.go .

RUN mage installdeps

COPY internal internal

ENTRYPOINT ["mage", "test"]
