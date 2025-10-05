FRONTEND=frontend-service
BROKER=broker-service
AUTH=auth-service

## up: starts all containers without forcing build
up:
	@echo "Starting Docker services..."
	docker-compose up -d
	@echo "Services started!"

## up_build: forces rebuild and restarts
up_build:
	@echo "Stopping any running containers..."
	docker-compose down
	@echo "Rebuilding and starting services..."
	docker-compose up --build -d
	@echo "Services built and started!"

## down: stop containers
down:
	@echo "Stopping services..."
	docker-compose down
	@echo "Done!"

## logs: show logs from all services
logs:
	docker-compose logs -f

## ps: show running services
ps:
	docker-compose ps

## restart: restart all services
restart: down up

## restart_front: restart only the frontend
restart_front:
	docker-compose restart $(FRONTEND)

## restart_broker: restart only the broker
restart_broker:
	docker-compose restart $(BROKER)

## restart_auth: restart only the auth
restart_auth:
	docker-compose restart $(AUTH)