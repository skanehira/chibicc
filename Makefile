CFLAGS=-std=c11 -g -static

test:
	./test.sh

clean: clean
	rm -rf chibicc *.o *~ tmp*

.PHONY: test clean
