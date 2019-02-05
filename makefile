BUILD_FLAGS = -tags netgo -a

deps:
	go get -u github.com/lib/pq
	go get -u github.com/pressly/goose/cmd/goose
	go get -u github.com/gorilla/mux
	go get -u github.com/google/uuid
	go get -u github.com/stretchr/testify/assert
	go get -u github.com/stretchr/testify/require

build:
	goose -dir "migrations" postgres "user=${DBUSER} password=${DBPASS} dbname=postgres_db sslmode=disable" up
	goose -dir "migrations" postgres "user=${DBUSER} password=${DBPASS} dbname=postgres_test_db sslmode=disable" up
	go build $(BUILD_FLAGS) ./...

install: build
	go install $(BUILD_FLAGS)

test:
	go clean -testcache
	go test ./... -v -p 1

runserver: install
	./bin/sample-go-api

fmt:
	gofmt -w ./
