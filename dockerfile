# Dockerfile
FROM golang:1.20 as builder

WORKDIR /app
COPY . .

RUN go build -o go-router .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/go-router .

CMD ["./go-router"]
