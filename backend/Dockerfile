FROM golang:1.23.6-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* .

RUN go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]