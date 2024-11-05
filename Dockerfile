FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .
COPY internal/infrastructure/database/migrations /migrations

RUN apk add --no-cache postgresql-client

ENV DB_HOST=db
ENV DB_USER=postgres
ENV DB_PASSWORD=1234
ENV DB_NAME=urlshortener
ENV DB_PORT=5432

CMD ["sh", "-c", "until psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -W -c 'select 1'; do sleep 1; done && psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -W -f /migrations/00001_create_urls_table.down.sql && ./main"]

EXPOSE 8000
