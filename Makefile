.PHONY: clean all build run test-asm test-vmx 

usage:
	@echo "usage: make [all|build|clean|test]"

clean:
	rm -f n2t-asm n2t-vmx *.out *.asm *.hack
	go clean -testcache

all build: n2t-asm n2t-vmx

n2t-asm: services/asm/*.go libs/n2t/*.go
	go build -o n2t-asm services/asm/*.go

n2t-vmx: services/vmx/*.go
	go build -o n2t-vmx services/vmx/*.go

lint:
	go fmt services/asm/*.go
	go fmt services/vmx/*.go

test: n2t-asm n2t-vmx
	go test -v ./services/asm

SimpleAdd.asm: n2t-asm servicesjdpkg/vmx/testdata/SimpleAdd/SimpleAdd.vm
	cp services/vmx/testdata/SimpleAdd/SimpleAdd.vm .
	./n2t-vmx SimpleAdd.vm
