.PHONY: clean all build run

usage:
	@echo "usage: make [all|build|clean|test]"

clean:
	rm -f testfiles/*/*/*.asm
	rm -rf project8

realclean:
	rm -f VMTranslator *.zip
	go clean -testcache

all build: VMTranslator

VMTranslator: *.go
	go build -o VMTranslator *.go

lint:
	go fmt *.go
	go vet *.go

test: VMTranslator
	./VMTranslator testfiles/ProgramFlow/BasicLoop    

zip project8.zip:
	mkdir -p project8
	cp -a *.go project8/
	echo "go" > project8/lang.txt
	echo "debug" >> project8/lang.txt
	zip -r project8.zip project8/