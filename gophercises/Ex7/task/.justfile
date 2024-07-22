default: help

bin := "./bin/task"

set quiet
# build to bin/task
build:
    go build -o {{bin}} main.go

# build and show help
help: build
    {{bin}}

add +ARGS: build
    {{bin}} add {{ARGS}}

do +ARGS: build
    {{bin}} do {{ARGS}}

list: build
    {{bin}} list

# clean and reset
[no-quiet]
clean:
    rm -rf bin/*
