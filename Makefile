#!make
include .env

GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.DEFAULT_GOAL := run

.PHONY: all build run clean test coverage create_db drop_db migrate_up migrate_down sqlc develop deps_reset tidy deps_upgrade deps_cleancache help

all: help

## Build;
build: ## Build your project and put the output binary in bin/
	mkdir -p bin
	GO111MODULE=on GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY}-darwin ./cmd/gohtmx
	GO111MODULE=on GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY}-linux ./cmd/gohtmx
	GO111MODULE=on GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY}-windows ./cmd/gohtmx

run: build ## Build and run your project
	./bin/${BINARY}-linux
	
clean: ## Remove build related files
	go clean
	rm -f ./bin/${BINARY}-*

## Test:
test: ## Run the tests of the project
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/jstemmer/go-junit-report
	$(eval OUTPUT_OPTIONS = | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml)
endif
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

coverage: ## Run the tests of the project and export the coverage
	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov
ifeq ($(EXPORT_RESULT), true)
	GO111MODULE=off go get -u github.com/AlekSi/gocov-xml
	GO111MODULE=off go get -u github.com/axw/gocov/gocov
	gocov convert profile.cov | gocov-xml > coverage.xml
endif

## Database:
create_db: ## Create the database
	docker exec -it ${DB_NAME}-db createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

drop_db: ## Drop the database
	docker exec -it ${DB_NAME}-db dropdb ${DB_NAME}

migrate_up: ## Run the up migrations
	goose -dir sql/schema postgres "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

migrate_down: ## Run the down migrations
	goose -dir sql/schema postgres "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down

sqlc: ## Generate the sqlc code
	sqlc generate

## Development:
build_tmp:
	go build -o ./tmp/gohtmx ./cmd/gohtmx

develop: build_tmp ## Run the project in development mode
	echo "Starting docker environment"
	docker compose -d --build

develop_stop: ## Stop the docker environment
	docker compose down

## Deps:
deps_reset: ## Reset the dependencies
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy: ## Tidy the dependencies
	go mod tidy
	go mod vendor

deps_upgrade: ## Upgrade the dependencies
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps_cleancache: ## Clean the dependencies cache
	go clean -modcache

## Help:
help: ## Display this help screen
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)