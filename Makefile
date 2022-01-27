include .env
export

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up
migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down -all
