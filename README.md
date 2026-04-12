# Tasks API

REST API для управления задачами на Go с Clean Architecture и DDD.

## Стек

- **Go** 1.21+ | **PostgreSQL** 18 | **Docker**
- **Chi** (роутинг) | **Zerolog** (логи) | **pgx** (БД)
- **Validator** (валидация) | **golang-migrate** (миграции)

## Структура
```
.
├── cmd/app/main.go           # Точка входа
├── config/                   # Конфигурация
├── internal/
│   ├── domain/              # Сущности и интерфейсы
│   ├── usecase/             # Бизнес-логика
│   ├── dto/                 # Request/Response модели
│   ├── adapter/postgres/    # Репозитории
│   └── controller/http/     # HTTP handlers
├── pkg/
│   ├── logger/              # Zerolog
│   ├── transaction/         # Транзакции
│   └── httpserver/          # HTTP сервер
├── migrations/              # SQL миграции
└── docker-compose.yaml
```

## Makefile команды

```bash
make run              # Собрать и запустить всё
make env-down         # Остановить
make env-cleanup      # Удалить контейнеры и данные

make migrate-up       # Применить миграции
make migrate-down     # Откатить миграцию

make build            # Собрать образы
make build-app        # Собрать только приложение
```

## Endpoints
| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| POST | /api/v1/tasks | Создать задачу |
| GET | /api/v1/tasks?sort={sort}&order={order} | Список задач |
| GET | /api/v1/tasks/{id} | Получить задачу |
| PATCH | /api/v1/tasks/{id} | Обновить задачу |
| DELETE | /api/v1/tasks/{id} | Удалить (soft delete) |

```bash
# Создать
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить хлеб","description":"Срочно"}'

# Получить список с сортировкой
curl "http://localhost:8080/api/v1/tasks?sort=created_at&order=desc"

# Обновить
curl -X PATCH http://localhost:8080/api/v1/tasks/{id} \
  -H "Content-Type: application/json" \
  -d '{"completed":true}'

# Удалить
curl -X DELETE http://localhost:8080/api/v1/tasks/{id}
```
## БД
```bash
# Доступ к БД
docker compose exec tasks-postgres psql -U test-user -d test-db
```
## .env
```env
# === Database ===
POSTGRES_USER=test-user
POSTGRES_PASSWORD=test-postgres-password
POSTGRES_DB=test-db
POSTGRES_PORT=5432
POSTGRES_HOST=tasks-postgres

# === Application ===
APP_NAME=tasks
APP_VERSION=1.0.0

# === HTTP Server ===
HTTP_HOST=0.0.0.0
HTTP_PORT=8080
HTTP_SHUTDOWN_TIMEOUT=30s

# === Logger ===
LOGGER_LEVEL=debug
LOGGER_PRETTY=true
# LOGGER_FOLDER=./out // not implemented yet
```