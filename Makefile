hello:
	@echo "Hello Golang!"

mycprog: hello.c
	cc -o $@ $<

clean:
	rm -f mycprog myprog

