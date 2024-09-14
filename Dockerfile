FROM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o hexaprogress /app/main.go

FROM alpine:latest

WORKDIR /app

COPY /source /app/source
COPY --from=builder /app/hexaprogress /app/hexaprogress

CMD ["./hexaprogress"]