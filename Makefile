usage:
	@echo "usage: make [all|build|clean|test]"

all: build

n2t build:
	go build -o n2t main.go

clean:
	rm -rf n2t

test:
	go test -v ./...
