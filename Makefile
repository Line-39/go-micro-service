#!make
include .env
export

config:
	@echo -e 'Current config:\n-Name:\t\t${NAME}\n-Version:\t${VERS}\n-Address:\t${ADDR}\n'

run:
	@go run main.go >> log

build:
	@go build .

all:
	@$(MAKE) build
	@./go-microservice >> log