FROM golang:1.24.2-bookworm AS builder

WORKDIR /compiler

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o webconverter .

FROM alpine:3.21

RUN apk add --no-cache chromium

WORKDIR /app

COPY --from=builder /compiler/webconverter .

EXPOSE 3000

CMD ["./webconverter"]