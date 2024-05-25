# Make shell
default:
	devbox shell

launch:
# Install dependencies
	(cd ./backend && npm ci)
	(cd ./frontend && npm ci)
# Bring up the docker containers
	make down
	make up
# Migrate and seed database
	make seed-db
	@echo "----------------------------------------"
	@echo "Welcome to the Daytwo Dev Environment."
	@echo "What will you build today?"
	@echo "----------------------------------------"

# Setup docker containers
up:
	docker compose -f "compose.yml" up -d --build --wait

# Teardown docker containers
down:
	docker compose -f "compose.yml" down
	(echo "y" | docker volume prune)
	clear

# Migrate and seed database
seed-db:
	(cd ./backend && npm run db:migrate-dev)
