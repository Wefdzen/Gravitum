# Этап сборки
FROM golang:1.23.1-alpine AS builder

# Установка git
RUN apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копирование файлов зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование всего проекта
COPY . .

# Сборка приложения
RUN go build -o app ./cmd/app/main.go

# Финальный образ
FROM alpine:latest

# Установка сертификатов
RUN apk --no-cache add ca-certificates

# Установка рабочей директории
WORKDIR /app

# Копирование всего проекта из этапа сборки
COPY --from=builder /app .

# Открытие порта
EXPOSE 8080

# Запуск приложения
CMD ["./app"]
