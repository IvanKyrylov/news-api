# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
RESET  := $(shell tput -Txterm sgr0)
DATE = $(shell date -u '+%Y-%m-%dT%H:%M:%S')
VERSION ?= $(shell git describe --tags --abbrev=7)

define colored
	@echo '${GREEN}$1${RESET}'
endef

test:
	go test -race -coverprofile=cover.out -failfast -timeout 120s -parallel 1 ./... && \
		go tool cover -func cover.out && \
		if [ $$? -eq 0 ]; then \
			coverage=$$(go tool cover -func=cover.out |  tail -1|grep -Eo "[0-9]+" | head -1); \
			if [ $$coverage -lt 25 ]; then \
				echo COVERAGE is less then 25 && exit 1; \
			fi \
	  	fi

run:
	docker-compose up -d

build:
	go build -v .

lint:
	golangci-lint run -v ./...