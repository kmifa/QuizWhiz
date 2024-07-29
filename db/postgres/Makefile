.PHONY: migrate/up
migrate/up:
	@goose -dir="./migration" postgres "user=$(DB_USER_NAME) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=127.0.0.1 port=$(DB_TCP_PORT) sslmode=disable" up

.PHONY: migrate/down
migrate/down:
	@goose -dir="./migration" postgres "user=$(DB_USER_NAME) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=127.0.0.1 port=$(DB_TCP_PORT) sslmode=disable" down

.PHONY: test/up
test/up:
	@goose -dir="./test" postgres "user=$(DB_USER_NAME) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=127.0.0.1 port=5555 sslmode=disable" up
