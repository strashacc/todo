FROM golang:1.24 AS builder

WORKDIR /app

# Установить protoc и плагины
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Сначала скопировать только go.mod и go.sum для кэширования слоёв
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копировать остальной исходный код
COPY . .

# Сгенерировать Go-код из proto
RUN PATH="$PATH:$(go env GOPATH)/bin" \
    && protoc --go_out=. --go-grpc_out=. proto/notification.proto

ENV GOOS=linux
ENV GOARCH=amd64

# Собрать приложение
RUN go build -o notification-service main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=builder /app/notification-service .
COPY --from=builder /app/.env .

# Если хочешь явно открыть порты (для info, не обязательно):
# EXPOSE 50051 8082

CMD ["./notification-service"]
