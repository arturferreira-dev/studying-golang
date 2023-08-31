FROM golang:alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init github.com/arturferreira-dev/studying-golang && \
    go mod tidy && \
    go build -o app -ldflags "-s -w" main.go
FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT ["./app"]