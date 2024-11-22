.PHONY: clean all build run test-asm test-vmx 

usage:
	@echo "usage: make [all|build|clean|test]"

clean:
	rm -f vmx *.out *.asm *.hack
	rm -rf project7
	go clean -testcache

all build: vmx

vmx: src/*.go
	go build -o vmx src/*.go

lint:
	go fmt src/*.go

zip project7.zip:
	mkdir -p project7
	cp -a src/*.go src/go.* project7/
	echo "go" > project7/lang.txt
	echo "debug" >> project7/lang.txt
	zip -r project7.zip project7/
