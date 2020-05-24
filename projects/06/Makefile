usage:
	@echo "usage: make [zip|build|all|clean|asm]"

clean:
	rm -f *.hack *.tmp *.zip asm

zip project06.zip: 
	zip -j project06.zip main.go *.hack
	unzip -l project06.zip

all: Add.hack

build asm: main.go
	go build

Add.hack: asm etc/Add.asm
	./asm etc/Add.asm > Add.hack

run: asm
	./asm etc/Add.asm

test: asm
	go test -v
	./asm etc/Add.asm > Add.tmp; diff Add.tmp etc/Add.hack
