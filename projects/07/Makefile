.PHONY: clean all build run test-asm test-vmx 

usage:
	@echo "usage: make [all|build|clean|test]"

clean:
	rm -f vmx *.out *.asm *.hack
	go clean -testcache

all build: vmx

vmx: src/*.go
	go build -o vmx src/*.go

lint:
	go fmt src/*.go
