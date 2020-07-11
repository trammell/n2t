
clean:
	rm -f *.jar

run: vmtrans.jar
	java -jar vmtrans.jar StackArithmetic/SimpleAdd/SimpleAdd.vm

vmtrans.jar: vmtrans.kt
	kotlinc vmtrans.kt -include-runtime -d vmtrans.jar
