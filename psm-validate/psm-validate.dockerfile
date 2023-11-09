FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o ./cmd/api/main

EXPOSE 8080

CMD ["./cmd/api/main"]