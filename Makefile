# Make shell
default:
	devbox shell

launch:
# Install dependencies
	(cd ./backend && go get ./... && go mod tidy && go generate ./...)
# Bring up the docker containers
	make down
	make up
# Migrate and seed database
	(cd ./backend && go run ./cmd/goosecli up)
	@echo "---------------------------------------------------"
	@echo "Welcome to the Transfers System Dev Environment."
	@echo "What will you build today?"
	@echo "---------------------------------------------------"

# Setup docker containers
up:
	docker compose -f "compose.yml" up -d --build --wait

# Teardown docker containers
down:
	docker compose -f "compose.yml" down
	(echo "y" | docker volume prune)
	clear

run-dev:
	(cd backend && air)

generate-docs:
	(cd backend && swag init)
	cp -r ./backend/docs ./api_docs/docs
	(cd api_docs && npm ci && npm run dev)
