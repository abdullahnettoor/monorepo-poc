.PHONY: init up down

init:
	@echo "Installing Go dependencies..."
	@go work sync
	@echo "Installing Node dependencies..."
	@npm install

up:
	@echo "Starting infrastructure..."
	@docker-compose up -d

down:
	@echo "Stopping infrastructure..."
	@docker-compose down
