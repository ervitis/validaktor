tests:
	go test -race -v ./...

lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

check: lint tests

cover:
	go test -race -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out
	cat cover.out >> coverage.txt