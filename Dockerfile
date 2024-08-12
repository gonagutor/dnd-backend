# Build stage
FROM golang:1.22.6 AS build
WORKDIR /dnd
COPY . .
RUN go mod download
RUN go build -o dnd-backend

# Production stage
FROM debian:latest AS production
WORKDIR /dnd
COPY --from=build /dnd/dnd-backend ./
COPY --from=build /dnd/templates ./templates
COPY --from=build /dnd/static ./static
CMD ./dnd-backend