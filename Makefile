.PHONY: build clean deploy

build:
	go mod tidy
	env GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/server/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose --aws-s3-accelerate
