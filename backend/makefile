.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run:
	# @go run ./src/
	go run github.com/cortesi/modd/cmd/modd --file=modd-run.conf

build:
	env GOOS=linux GOARCH=amd64 go build -o poker ./src/

publish: build
	rsync --rsync-path="sudo /usr/bin/rsync" ./poker matrix:/opt/poker
