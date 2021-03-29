run:
	SERVER_PORT=8080 go run github.com/ehabterra/money-transfers/cmd/server/

build:
	go build -o bin/money-transfers github.com/ehabterra/money-transfers/cmd/server

test:
	go test -cover -covermode=count ./internal/multiplexer/ ./internal/services/

coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...

curl-add-account:
	curl --header "Content-Type: application/json" \
      --request POST \
      --data '{"account_no":"100","balance":1000}' \
      http://localhost:8080/accounts