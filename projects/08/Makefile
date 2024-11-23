.PHONY: clean all build run test-asm test-vmx 

usage:
	@echo "usage: make [all|build|clean|test]"

clean:
	rm -f VMTranslator *.asm *.hack *.out *.zip
	rm -rf project8
	go clean -testcache

all build: vmx

VMTranslator: src/*.go
	go build -o VMTranslator src/*.go

lint:
	go fmt src/*.go
	go vet src/*.go

zip project8.zip:
	mkdir -p project8
	cp -a src/*.go src/go.* project8/
	echo "go" > project8/lang.txt
	echo "debug" >> project8/lang.txt
	zip -r project8.zip project8/
