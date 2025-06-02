
<h1 align="center">ToDo App — REST API на Go</h1>

<p align="center">
  <b>Go • PostgreSQL • Docker • JWT • REST API</b>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Backend-Go-blue?logo=go">
  <img src="https://img.shields.io/badge/Database-PostgreSQL-blue?logo=postgresql">
  <img src="https://img.shields.io/badge/Container-Docker-blue?logo=docker">
  <img src="https://img.shields.io/badge/Auth-JWT-orange">
</p>

---

## 📌 Описание проекта

REST API сервис для управления списками задач (ToDo), разработанный с использованием чистой архитектуры.

**Основные возможности:**
- Регистрация и авторизация пользователей через JWT
- Создание списков задач
- Добавление, обновление и удаление задач
- Защищённый доступ к данным пользователя
- Хэширование паролей для безопасного хранения

---

## 🔧 Используемый стек

- Go (Golang)
- PostgreSQL
- Docker / Docker Compose
- JWT авторизация
- Чистая архитектура (Clean Architecture)
- GitLab CI (для сборки)
- GolangCI-Lint

---

## 📂 Структура проекта

- `cmd/todo/main.go` — точка входа в приложение
- `internal/handler/` — обработка HTTP-запросов
- `internal/service/` — бизнес-логика
- `internal/repository/` — работа с БД
- `internal/models/` — описания сущностей (User, TodoList, TodoItem)
- `internal/config/` — конфигурация приложения
- `internal/server/` — запуск HTTP-сервера
- `db/migrations/` — миграции базы данных

---

## 🚀 Развёртывание проекта

### Сборка Docker-образа:

```bash
docker build -t todo-app .
```

### Запуск через Docker Compose:

```bash
docker-compose up --build
```

### Запуск локально (без Docker):

```bash
go run cmd/todo/main.go
```

⚠ Перед этим необходимо создать базу данных и применить миграции из `db/migrations/`

---

## ⚙️ Миграции базы данных

SQL-файлы для создания таблиц находятся в папке:

```
db/migrations/
```

---

## 📄 Документация API

API реализовано в стиле REST:

- `POST /auth/sign-up` — регистрация
- `POST /auth/sign-in` — авторизация
- `POST /api/lists` — создание списка
- `POST /api/items` — создание задачи
- и другие

---

## 📬 Контакты разработчика

- Автор: Молоканов Алексей Александрович
- Telegram: [@bobr_lord](https://t.me/bobr_lord)

---
