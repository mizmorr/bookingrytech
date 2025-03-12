run:
	cd cmd; go run main.go

test:
	go test -v -cover ./...

swag:
	swag init -g cmd/main.go -o docs
