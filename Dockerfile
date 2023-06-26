FROM golang:alpine

WORKDIR /revosearch
COPY . .
RUN go build -o /revosearch/backend main.go