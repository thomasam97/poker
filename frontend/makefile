
.PHONY: help, build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run:
	@yarn install
	@yarn dev

build:
	@yarn install
	@yarn build

publish: #build
	rsync -v -r --rsync-path="sudo /usr/bin/rsync" ./build/* matrix:/var/www/poker/

preview:
	http-server ./build -d -i 