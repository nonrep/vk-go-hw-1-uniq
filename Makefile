# Makefile

BINARY=uniq

all: build

build:
	go build -o $(BINARY) .

.PHONY: all build