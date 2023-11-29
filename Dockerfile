FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /migrate ./db/migrate/main.go

RUN go build -o /healthcare-api

RUN find . -name "*.go" -type f -delete

ENTRYPOINT ["./entrypoint.sh"]
