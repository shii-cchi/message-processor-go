# Message Processor Go

## Обзор

Микросервис, `message-processor-service`, который:
- принимает сообщения через HTTP API
- сохраняет их в PostgreSQL
- отправляет сообщения в Kafka для дальнейшей обработки, при этом обработанные сообщения помечаются
- предоставляет API для получения статистики по обработанным сообщениям

В комплект входит тестовый сервис `test-service`, предназначенный для эмуляции обработки сообщений.

## API Эндпоинты

### Прием сообщений

- **Эндпоинт:** `/messages`
- **Метод:** `POST`
- **Описание:** Этот эндпоинт принимает сообщения для обработки.
- **Пример запроса:**

```json
{
  "content": "Это тестовое сообщение"
}
```

Проверить тут:

```
http://45.138.74.112:8888/messages
```

### Получение статистики

- **Эндпоинт:** `/messages/stats`
- **Метод:** `GET`
- **Описание:** Этот эндпоинт предоставляет статистику по обработанным сообщениям.

Проверить тут:

```
http://45.138.74.112:8888/messages/stats
```

## Переменные окружения

Пример .env файла:

   ```
    PORT=8888
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_HOST=postgres
    DB_PORT=5432
    DB_NAME=messages
    KAFKA_BROKER=kafka:9092
    KAFKA_PORT=9092
    KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=INTERNAL:PLAINTEXT,INTERNAL_SAME_HOST:PLAINTEXT
    KAFKA_LISTENERS=INTERNAL://:9092
    KAFKA_ADVERTISED_LISTENERS=INTERNAL://kafka:9092
    KAFKA_INTER_BROKER_LISTENER_NAME=INTERNAL
    KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR='1'
    KAFKA_MIN_INSYNC_REPLICAS='1'
    KAFKA_CREATE_TOPICS="processed-messages:1:1,unprocessed-messages:1:1"
   ```

## Требования

- Docker

## Начало работы

Склонируйте репозиторий, создайте .env файл и из папки `message-processor-go` запустите:

   ```
   docker network create message_network && docker-compose up -d
   ```