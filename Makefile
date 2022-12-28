:PHONY build
build:
	go build -v ./cmd/apiserver

:PHONY test
test:
	go test -v -race -timeout 3s ./...

:PHONY up_db
up_db:
	migrate -path migrations -database "postgres://localhost/restapi_dev?user=postgres&password=myPassword&sslmode=disable" up

:PHONY down_db
down_db:
	migrate -path migrations -database "postgres://localhost/restapi_dev?user=postgres&password=myPassword&sslmode=disable" down

.DEFAULT_GOAL := build