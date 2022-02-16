usage:
	@echo "usage: make [all|build|clean|test]"

all: build

n2t build:
	go build -o n2t main.go

clean:
	rm -rf n2t
	go clean -testcache ./...

# test building blocks first, then work up to integration tests
test:
	go clean -testcache ./...
	go test -v pkg/asm/{source,symboltable}_test.go
	go test -v pkg/asm/{a,c,label}_test.go
	go test -v ./...

coverage:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
