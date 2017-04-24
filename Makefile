default: server

install:
	@cp .env.sample .env

server:
	@env $$(cat .env) go run main.go
