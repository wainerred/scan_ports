#
FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server

#
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/server .
COPY servers.json .

EXPOSE 8080

CMD ["./server"]
