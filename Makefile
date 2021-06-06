NAME=gbc-banking
OS ?= linux

.PHONY: dev-up
dev-up:
	@echo ">>>>> Starting server application..."
	docker-compose up --build -d

.PHONY: dev-down
dev-down:
	@echo ">>>>> Shutting application..."
	docker-compose down

.PHONY: test
test:
	@echo "==> Running Tests..."
	go test -v ./...
