.DEFAULT_GOAL = bne

build:
	go build -o bin/main main.go

execute:
	bin/main

bne: 
	$(MAKE) build
	$(MAKE) execute