tidy:
	go mod tidy

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

test:
	go test

coverage:
	go test . -count=1 -coverprofile /tmp/cover.out
	go tool cover -html=/tmp/cover.out