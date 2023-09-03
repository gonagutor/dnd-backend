FROM golang:alpine

WORKDIR /dnd
COPY . .
RUN go build -o dnd-backend

RUN ./dnd-backend