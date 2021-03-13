run:
	SERVER_PORT=8080 go run github.com/ehabterra/money-transfers/cmd/server/

build:
	go build -o bin/money-transfers github.com/ehabterra/money-transfers/cmd/server

test:
	go test -cover ./internal/multiplexer/ ./internal/services/


curl-add-account:
	curl --header "Content-Type: application/json" \
      --request POST \
      --data '{"account_no":"100","balance":1000}' \
      http://localhost:8080/accounts