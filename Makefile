#!make
include .env

.DEFAULT_GOAL := run

.PHONY: build run clean create_db drop_db migrate_up migrate_down sqlc build_tmp develop deps_reset tidy deps_upgrade deps_cleancache

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${API_BINARY}-darwin ./cmd/gohtmx
	GOARCH=amd64 GOOS=linux go build -o ./bin/${API_BINARY}-linux ./cmd/gohtmx
	GOARCH=amd64 GOOS=windows go build -o ./bin/${API_BINARY}-windows ./cmd/gohtmx

run: build
	./bin/${API_BINARY}-linux
	
clean:
	go clean
	rm -f ./bin/${API_BINARY}-*

create_db:
	docker exec -it ${DB_NAME}-db createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

drop_db:
	docker exec -it ${DB_NAME}-db dropdb ${DB_NAME}

migrate_up:
	goose -dir sql/schema postgres "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

migrate_down:
	goose -dir sql/schema postgres "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down

sqlc:
	sqlc generate

build_tmp:
	go build -o ./tmp/gohtmx ./cmd/gohtmx

develop: build_tmp
	echo "Starting docker environment"
	docker compose -f docker-compose.yml up -d --build

deps_reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps_upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps_cleancache:
	go clean -modcache
