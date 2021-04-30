CFLAGS=-std=c11 -g -static

.PHONY: build
build:
	go build -o chibicc

.PHONY: build
test: build
	./test.sh

.PHONY: clean
clean: clean
	rm -rf chibicc *.o *~ tmp*

.PHONY: test clean
