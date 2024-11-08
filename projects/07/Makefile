.PHONY: clean all build run test-asm test-vmx 

usage:
	@echo "usage: make [all|build|clean|test]"

clean:
	rm -f vmx *.out *.asm *.hack
	go clean -testcache

all build: vmx

vmx: *.go
	go build -o vmx *.go

lint:
	go fmt *.go

SimpleAdd.asm: n2t-asm servicesjdpkg/vmx/testdata/SimpleAdd/SimpleAdd.vm
	cp services/vmx/testdata/SimpleAdd/SimpleAdd.vm .
	./n2t-vmx SimpleAdd.vm
