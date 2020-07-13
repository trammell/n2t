
clean:
	rm -f *.jar *.asm

run: vmtrans.jar
	java -jar vmtrans.jar StackArithmetic/SimpleAdd/SimpleAdd.vm | tee simpleadd.asm
	#java -jar vmtrans.jar StackArithmetic/StackTest | tee stacktest.asm

vmtrans.jar: vmtrans.kt
	kotlinc vmtrans.kt -include-runtime -d vmtrans.jar
