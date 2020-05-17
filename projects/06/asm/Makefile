
usage:
	@echo "usage: make [build]"

clean:
	rm -f asm

build asm: main.go
	go build

run: asm
	./asm file.txt
