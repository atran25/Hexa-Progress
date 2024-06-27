FROM golang:alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o hexaprogress /app/cmd/hexaprogress/

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/hexaprogress /app/hexaprogress

CMD ["./hexaprogress"]