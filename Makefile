usage:
	@echo "usage: make [all|build|clean|test]"

all: build

n2t build: cmd/*.go pkg/asm/*.go
	go build -o n2t main.go

clean:
	rm -f n2t coverage.out
	go clean -testcache ./...

test:
	go test -v ./...

coverage:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...

run: n2t
	./n2t asm pkg/asm/testdata/Max.asm
