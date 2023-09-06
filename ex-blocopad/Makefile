GOPACKAGES=$(shell go list ./... | egrep -v "vendor|mock")

.PHONY: run install test lint

GO = go
GOLANGCILINT = golangci-lint

run:
	$(GO) run cmd/main.go

run/api: run

clean:
	rm -rf vendor

install: clean tools
	$(GO) mod vendor && $(GO) mod tidy

test: deps/restart
	$(GO) test -tags="all" -count=1 -v $(GOPACKAGES)

test/unit:
	ENVIRONMENT=test $(GO) test -short -tags="unit" -count=1 -v $(GOPACKAGES)

test/integration: deps/restart
	ENVIRONMENT=test $(GO) test -run Integration -tags="integration" -count=1 -v $(GOPACKAGES)

lint:
	$(GOLANGCILINT) run

swagger:
	swag init -generalInfo server/server.go -output ./docs/swagger

clean-coverage:
	rm -rf cover.out

coverage: clean-coverage deps/restart
	ENVIRONMENT=test $(GO) test -tags="all" -covermode="count" -coverprofile="cover.out" ./...

coverage-html: coverage
	$(GO) tool cover -html=cover.out

deps/up:
	docker compose up --build -d

deps/down:
	docker compose down -v

deps/restart: deps/down deps/up

db/create-migration:
	$(_goose_) create $(MIGRATION_NAME) sql

db/migrate-up:
	$(_goose_) up

db/migrate-down:
	$(_goose_) down

db/migration-status:
	$(_goose_) status