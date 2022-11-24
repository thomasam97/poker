
.PHONY: help, build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: install
	@pnpm dev

build: install
	@pnpm build

publish: build
	zip -r build.zip ./build/*
	scp build.zip poker:~/opt/poker

install: ## Install dependencies
	@pnpm install

preview:
	http-server ./build -d -i 