.DEFAULT_GOAL := bne

build:
	go build -o bin/main main.go

execute:
	./bin/main

run:
	go run main.go

bne:
	$(MAKE) build
	$(MAKE) execute