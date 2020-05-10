NAME = "serv"

h help:
	@echo 	"h help - Справка"
	@echo	"build - Сборка проекта"
	@echo	"run - Сборка и запуск"
.PHONY: h help

build: swag
	go build -o ${NAME}
.PHONY: build

run: build
	./${NAME}
.PHONY: run

test:
	go test ./... -v
.PHONY: test

swag: 
	swag init
.PHONY: swag

dev: swag run
.PHONY: dev