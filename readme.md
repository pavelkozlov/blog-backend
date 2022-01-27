![example workflow](https://github.com/pavelkozlov/blog-backend/actions/workflows/golangci-lint.yml/badge.svg)
# Цель проекта

Реализовать аналог статического блога

## Миграции
Установить migrate CLI
```shell
brew install golang-migrate
```
Добавить файл миграции
```shell
migrate create -ext sql -dir migrations -seq init
```
Миграция
```shell
make migrate-up
make migrate-down
```