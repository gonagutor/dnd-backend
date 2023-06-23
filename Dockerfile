FROM golang:alpine

WORKDIR /flat-searcher
COPY . .
RUN go build -o /flat-searcher/backend main.go