FROM golang:1.21-alpine3.17

LABEL mantainer="danieeeld2@correo.ugr.es" \
      version="5.0.0"

RUN addgroup -S testgroup && \
    adduser -S -u 1001 test -G testgroup

WORKDIR /user/test

RUN apk update && \
    apk add git && \
    go install github.com/magefile/mage@v1.15.0 && \
    chown -R test:testgroup /user/test

USER test

COPY go.mod go.sum magefile.go ./
RUN mage installdeps

COPY internal internal

ENTRYPOINT ["mage", "test"]