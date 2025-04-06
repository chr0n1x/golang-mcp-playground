default: entr

entr:
	find *.go | entr -r \
		bash -c \
		'gofmt -w . && docker rmi -f golang-mcp-playground-mcp && docker compose build'

dev:
	docker compose run --rm -ti mcp bash
