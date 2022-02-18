include .env
export

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up
migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down -all
test-coverage:
	go tool cover -html=coverage.out
lint:
	go generate ./...
	golangci-lint run ./... && \
	goimports -local "github.com/pavelkozlov/blog-backend" -w . && \
	go fmt ./... && \
	go mod tidy
	 go test -v -coverprofile=coverage.out ./...