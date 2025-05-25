FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o hoststats

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/hoststats /hoststats

EXPOSE 8080

ENTRYPOINT ["/hoststats"]
