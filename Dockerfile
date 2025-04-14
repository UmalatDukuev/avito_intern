FROM golang:1.23.0-alpine as builder

COPY . /avito_intern
WORKDIR /avito_intern

RUN go mod download
RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]