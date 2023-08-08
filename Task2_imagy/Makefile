build:
	@go build -o imagy main.go

help: build
	@./imagy

run: build
	@./imagy run

download: build
	@./imagy download

compose-up:
	@docker compose up -d