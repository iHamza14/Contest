#  we can use it by simply writing command like {make migrate-up}
include backend/.env

MIGRATIONS_PATH := backend/migrations
DATABASE_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.PHONY: migrate-up migrate-down migrate-version migrate-force migrate-create

migrate-up:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

migrate-down:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

migrate-version:
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

migrate-force:
	@test -n "$(VERSION)" || (echo "Usage: make migrate-force VERSION=1" && exit 1)
	@migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(VERSION)

migrate-create:
	@test -n "$(NAME)" || (echo "Usage: make migrate-create NAME=create_submissions" && exit 1)
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME)
