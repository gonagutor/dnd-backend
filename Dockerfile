# Build stage
FROM golang:1.20.4 as build
WORKDIR /dnd
COPY . .
RUN go mod download
RUN go build -o dnd-backend

# Production stage
FROM debian:latest as production
WORKDIR /dnd
COPY --from=build /dnd/dnd-backend ./
COPY --from=build /dnd/templates ./
CMD ./dnd-backend