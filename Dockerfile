FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM golang:1.23

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/schema/migrations /app/schema/migrations
COPY --from=builder /app/config /app/config
COPY --from=builder /app/entrypoint.sh /app/entrypoint.sh
COPY --from=builder /app/.env /app/.env

RUN chmod +x /app/entrypoint.sh
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENV DATABASE_URL=postgres://postgres:pass@db:5432/avito?sslmode=disable

CMD ["/app/entrypoint.sh"]
