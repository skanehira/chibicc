CFLAGS=-std=c11 -g -static

chibicc: chibicc
	./test.sh

clean: clean
	rm -rf chibicc *.o *~ tmp*

.PHONY: test clean
