# ai-stats-microservices

Два микросервиса на Go для сбора и отдачи статистики по вызываемым AI сервисам

## How to use

1. Создайте .env файл в корне репозитория
2. Заполните его данными ниже. Можно взять готовый пример из .env_example:
    ```bash
    POSTGRES_HOST=хост_сервиса_с_базой_данных
    POSTGRES_PORT=порт_сервиса_с_базой_данных
    POSTGRES_USER=пользователь_базы_данных
    POSTGRES_PASSWORD=пароль_пользователя_базы_данных
    POSTGRES_DB=имя_базы_данных
    GRPC_HOST=хост_сервиса_с_grpc
    GRPC_PORT=порт_сервиса_с_grpc
    HTTP_PORT=http_порт
    ```
3. Поднимите контейнеры, находясь в корне репозитория:
    ```bash
    docker-compose up -d --build
    ```
4. Запросы можно будет слать по ```http://localhost:8080```

## Примеры запросов

### POST /call

Содержимое запроса:
```json
{
    "user_id": 1,
    "service_id": 1
}
```

Адрес:
```http request
http://localhost:8080/call
```

### GET /calls

Адрес:
```http request
http://localhost:8080/calls?user_id=1&service_id=1&page=1&limit=10
```

### POST /service

Содержимое запроса:
```json
{
    "name": "service c",
    "description": null,
    "price": 300
}
```

Адрес:
```http request
http://localhost:8080/service
```