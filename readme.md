![example workflow](https://github.com/pavelkozlov/blog-backend/actions/workflows/golangci-lint.yml/badge.svg)
[![codecov](https://codecov.io/gh/pavelkozlov/blog-backend/branch/main/graph/badge.svg?token=7AIRY67CYH)](https://codecov.io/gh/pavelkozlov/blog-backend)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GoReportCard example](https://goreportcard.com/badge/github.com/pavelkozlov/blog-backend)](https://goreportcard.com/report/github.com/pavelkozlov/blog-backend)
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
