GO := go
AIR := $(shell which air 2>/dev/null)

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make run       - Air でホットリロード"
	@echo "  make migrate   - cmd/migrate/main.go を実行"
	@echo "  make build     - main.go をビルド"

.PHONY: run
run:
ifeq ($(AIR),)
	$(error "Air is not installed. Install it with: go install github.com/cosmtrek/air@v1.62.3")
endif
	buf generate && air

.PHONY: migrate
migrate:
	$(GO) run cmd/migrate/main.go

.PHONY: build
build:
	$(GO) build -o bin/app main.go
