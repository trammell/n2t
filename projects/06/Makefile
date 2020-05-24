usage:
	@echo "usage: make [zip|build|all|clean|asm]"

clean:
	rm -f *.hack *.tmp *.zip asm

zip project06.zip: 
	zip -j project06.zip *.go go.* *.hack
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
	./asm etc/Rect.asm > Rect.tmp; diff -u Rect.tmp etc/Rect.hack
	./asm etc/Max.asm > Max.tmp; diff Max.tmp etc/Max.hack
	./asm etc/Add.asm > Add.tmp; diff Add.tmp etc/Add.hack
	./asm etc/MaxL.asm > MaxL.tmp; diff MaxL.tmp etc/MaxL.hack
	./asm etc/PongL.asm > PongL.tmp; diff PongL.tmp etc/PongL.hack
	./asm etc/RectL.asm > RectL.tmp; diff RectL.tmp etc/RectL.hack
	./asm etc/Pong.asm > Pong.tmp; diff -u Pong.tmp etc/Pong.hack

