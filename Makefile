test:
	go test -coverprofile=cover.out ./...
		
run:
	docker-compose up -d

build:
	go build -o ./bin ./cmd/news-api

lint:
	golangci-lint run -v ./...