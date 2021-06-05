NAME=gbc-banking
OS ?= linux

.PHONY: dev-up
run-dev:
	@echo ">>>>> Starting server application..."
	docker-compose up --build -d

.PHONY: dev-down
run-dev:
	@echo ">>>>> Shutting application..."
	docker-compose down