# Message Processor Go

## Обзор

Микросервис, `message-processor-service`, который:
- принимает сообщения через HTTP API
- сохраняет их в PostgreSQL
- отправляет сообщения в Kafka для дальнейшей обработки, при этом обработанные сообщения помечаются
- предоставляет API для получения статистики по обработанным сообщениям

В комплект входит тестовый сервис `test-service`, предназначенный для эмуляции обработки сообщений.

## Требования

- Docker

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

## Начало работы

Склонируйте репозиторий и из папки `message-processor-go` запустите:

   ```
   docker-compose up -d
   ```

## Тестирование через интернет