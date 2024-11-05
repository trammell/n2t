
.PHONY: clean all build run test-asm test-vmx 

usage:
	@echo "usage: make [all|build|clean|test]"

clean:
	rm -f n2t n2t-asm n2t-vmx *.out *.asm *.hack
	go clean -testcache

all build: n2t-asm

n2t-asm: services/asm/*.go libs/n2t/*.go
	go build -o n2t-asm services/asm/*.go

lint:
	go fmt services/asm/main.go services/asm/parser.go

test: n2t-asm
	go test -v ./services/asm

coverage:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...

run: n2t
	./n2t asm pkg/asm/testdata/Max.asm

SimpleAdd.asm: n2t pkg/vmx/testdata/SimpleAdd/SimpleAdd.vm
	./n2t vmx pkg/vmx/testdata/SimpleAdd/SimpleAdd.vm
