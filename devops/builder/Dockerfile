ARG BASE_IMAGE=golang:1.19.0-alpine3.16

FROM ${BASE_IMAGE} AS builder

WORKDIR /go/src

ADD . /go/src/

RUN go mod download && go mod verify

RUN go build -v -o /stori ./cmd/*.go

FROM ${BASE_IMAGE} AS candidate

ARG NAME=stori

# Install nice to haves
RUN apk add --no-cache openssl ncurses-libs libstdc++ libgcc curl libressl htop nano

WORKDIR /go/src/

COPY --from=builder /$NAME .
COPY --from=builder /go/src/data /go/src/data
COPY --from=builder /go/src/migrations /go/src/migrations

EXPOSE 8080

ENTRYPOINT ["./stori"]
