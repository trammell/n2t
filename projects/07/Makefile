
clean:
	rm -f main.jar

run: main.jar
	java -jar main.jar

main.jar:
	kotlinc main.kt -include-runtime -d main.jar
