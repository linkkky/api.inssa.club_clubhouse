.PHONY: build clean deploy

build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
