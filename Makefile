include .env
export

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up
migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down -all
test-coverage:
	 go test -v -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
lint:
	golangci-lint run ./... &&