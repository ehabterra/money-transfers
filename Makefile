run:
	SERVER_PORT = 8080 go run github.com/ehabterra/money_transfer/cmd/server/

build:
	SERVER_PORT = 8080 go build -o bin/money_transfer github.com/ehabterra/money_transfer/cmd/server

test:
	go test -cover ./internal/multiplexer/ ./internal/services/


curl-add-account:
	curl --header "Content-Type: application/json" \
      --request POST \
      --data '{"account_no":"100","balance":1000}' \
      http://localhost:8080/accounts